package gen_container

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"net"
	"os"
)

type CreateBuildAndContainerOptions struct {
	Tag           string
	DockerFile    string
	ContainerName string
	Cmd           string
	Env           []string
	Tty           bool
	Mounts        map[string]string
}

func CreateBuildAndContainer(c *client.Client,
	imageBuildOpt types.ImageBuildOptions,
	config *container.Config,
	hostConfig *container.HostConfig,
	networkingConfig *network.NetworkingConfig,
	platform *ocispec.Platform,
	containerName string,
) (container.CreateResponse, error) {

	var (
		err error
		ctx = context.Background()
	)

	if len(imageBuildOpt.Tags) == 0 {
		return container.CreateResponse{}, fmt.Errorf("docker images tags not found")
	}

	buildContextTar, err := archive.TarWithOptions("./", &archive.TarOptions{})
	if err != nil {
		return container.CreateResponse{}, err
	}

	image, err := c.ImageBuild(ctx, buildContextTar, imageBuildOpt)

	if err != nil {
		return container.CreateResponse{}, err
	}

	defer image.Body.Close()
	CopyIO(os.Stderr, image.Body)

	runContainer, err := c.ContainerCreate(
		ctx,
		config,
		hostConfig,
		networkingConfig,
		platform,
		containerName,
	)
	return runContainer, nil
}

func CreateAttachContainer(c *client.Client, option CreateBuildAndContainerOptions) (types.HijackedResponse, error) {
	var (
		err error
		ctx = context.Background()
	)
	volumes, binds, err := convertMounts(option.Mounts)

	if err != nil {
		return types.HijackedResponse{}, err
	}

	runContainer, err := CreateBuildAndContainer(c,
		types.ImageBuildOptions{
			Tags:       []string{option.Tag},
			Dockerfile: option.DockerFile,
		},
		&container.Config{
			Image:        option.Tag,
			OpenStdin:    true,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          option.Tty,
			Cmd:          strslice.StrSlice{option.Cmd},
			Env:          option.Env,
			Volumes:      volumes,
		},
		&container.HostConfig{
			NetworkMode: "host",
			Binds:       binds,
		},
		nil,
		nil,
		option.ContainerName)

	if err != nil {
		return types.HijackedResponse{}, err
	}

	attachContainer, err := c.ContainerAttach(ctx, runContainer.ID, container.AttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return types.HijackedResponse{}, err
	}

	startCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = c.ContainerStart(startCtx, runContainer.ID, container.StartOptions{}); err != nil {
		return types.HijackedResponse{}, err
	}

	go closeContainerConn(c, runContainer.ID, attachContainer.Conn)
	return attachContainer, nil
}

func closeContainerConn(c *client.Client, ID string, conn net.Conn) {
	ctx := context.Background()
	statusCh, errCh := c.ContainerWait(ctx, ID, container.WaitConditionNotRunning)
	defer conn.Close()
	select {
	case <-statusCh:
	case <-errCh:
	}

}
