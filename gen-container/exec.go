package gen_container

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/term"
	"time"
)

func WaitExecExit(ctx context.Context, cli *client.Client, execID string, execContainer *types.HijackedResponse) {
	defer execContainer.Close()
	for {
		inspect, err := cli.ContainerExecInspect(ctx, execID)
		if err != nil {
			return
		}
		if !inspect.Running {
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func ExecContainer(c *client.Client, ID string, cmd string) (types.HijackedResponse, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	width, height, err := term.GetSize(0)
	if err != nil {
		return types.HijackedResponse{}, err
	}

	execContainer, err := c.ContainerExecCreate(ctx, ID, container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{cmd},
		ConsoleSize:  &[2]uint{uint(height), uint(width)},
	})

	if err != nil {
		return types.HijackedResponse{}, err
	}

	attachResponse, err := c.ContainerExecAttach(ctx, execContainer.ID, container.ExecAttachOptions{
		Tty: true,
	})

	waitCtx := context.Background()
	go WaitExecExit(waitCtx, c, execContainer.ID, &attachResponse)
	return attachResponse, nil
}
