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
	"github.com/moby/term"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func setContainerTTYSize(cli *client.Client, containerID string) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGWINCH)

	go func() {
		for range sig {
			size, err := term.GetWinsize(os.Stdin.Fd())
			if err != nil {
				return
			}

			err = cli.ContainerResize(context.Background(), containerID, container.ResizeOptions{
				Height: uint(size.Height),
				Width:  uint(size.Width),
			})
			if err != nil {
				logrus.Warnf("resize error %s", err)
			}
		}
	}()
	sig <- syscall.SIGWINCH
}

type CreateBuildAndContainerOptions struct {
	Tag           string
	DockerFile    string
	ContainerName string
	Cmd           string
	Env           []string
	Tty           bool
	Mounts        map[string]string
}

type Container struct {
	Option          CreateBuildAndContainerOptions
	CreateContainer container.CreateResponse
	AttachContainer types.HijackedResponse
	Client          *client.Client

	closed chan error
}

func NewContainer(c *client.Client, option CreateBuildAndContainerOptions) (*Container, error) {
	volumes, binds, err := convertMounts(option.Mounts)

	if err != nil {
		return nil, err
	}

	CreateContainer, err := CreateBuildAndContainer(c,
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
		return nil, err
	}

	ctx := context.Background()
	attachContainer, err := c.ContainerAttach(ctx, CreateContainer.ID, container.AttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return nil, err
	}

	startCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = c.ContainerStart(startCtx, CreateContainer.ID, container.StartOptions{}); err != nil {
		return nil, err
	}

	setContainerTTYSize(c, CreateContainer.ID)
	go closeContainerConn(c, CreateContainer.ID, attachContainer.Conn)

	return &Container{
		CreateContainer: CreateContainer,
		AttachContainer: attachContainer,
		Option:          option,
		Client:          c,
		closed:          make(chan error, 1),
	}, nil

}

func (co *Container) CloseAttachContainer() {
	co.AttachContainer.Close()
}

func (co *Container) Close() error {
	ctx := context.Background()
	err := co.Client.ContainerRemove(ctx, co.CreateContainer.ID, container.RemoveOptions{
		Force: true,
	})
	co.closed <- err
	return err
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

func (co *Container) Closed() error {
	return <-co.closed
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
