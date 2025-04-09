package gen_container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"os"
)

type CreateExecOptions struct {
	ImageName     string
	ContainerName string
	Cmd           string
}

// CreateExecContainer
// docker run -itd xxx cmd and docker exec -it xxx cmd
func CreateExecContainer(c *client.Client, dockerFile string, option CreateExecOptions) (types.HijackedResponse, error) {
	ctx := context.Background()

	buildContextTar, err := archive.TarWithOptions("./", &archive.TarOptions{})
	if err != nil {
		return types.HijackedResponse{}, nil
	}

	image, err := c.ImageBuild(ctx, buildContextTar, types.ImageBuildOptions{
		Tags: []string{
			option.ImageName,
		},
		Dockerfile: dockerFile,
	})

	if err != nil {
		return types.HijackedResponse{}, err
	}

	defer image.Body.Close()
	CopyIO(os.Stderr, image.Body)

	if err != nil {
		return types.HijackedResponse{}, err
	}

	createCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	runContainer, err := c.ContainerCreate(
		createCtx,
		&container.Config{
			Image:        option.ImageName,
			AttachStdin:  true,
			AttachStdout: true,
			AttachStderr: true,
			Tty:          true,
			Cmd:          strslice.StrSlice{option.Cmd},
		},
		&container.HostConfig{
			NetworkMode: "host",
		},
		nil,
		nil,
		option.ContainerName,
	)

	if err != nil {
		return types.HijackedResponse{}, err
	}

	startCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err = c.ContainerStart(startCtx, runContainer.ID, container.StartOptions{}); err != nil {
		return types.HijackedResponse{}, err
	}

	return ExecContainer(c, runContainer.ID, option.Cmd)
}
