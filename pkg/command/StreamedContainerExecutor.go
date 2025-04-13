package command

import (
	"bytes"
	"context"
	"dockgen/pkg/concurrency"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io"
	"net"
	"strings"
)

type StreamedContainerExecutor struct {
	delim  string
	stdout io.Reader
	stderr io.Reader
	stdin  net.Conn
}

func NewStreamedContainerExecutor(client *client.Client, ID string, conn net.Conn) (*StreamedContainerExecutor, error) {
	reader, err := client.ContainerLogs(context.Background(), ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})

	if err != nil {
		return nil, err
	}
	stdout, wStdout := io.Pipe()
	stderr, wStderr := io.Pipe()
	go stdcopy.StdCopy(wStdout, wStderr, reader)
	delim := "__CMD_DONE__"
	outReader := NewDelimitedReader(stdout, delim)
	errReader := NewDelimitedReader(stderr, delim)
	return &StreamedContainerExecutor{
		stdin:  conn,
		delim:  delim,
		stdout: outReader,
		stderr: errReader,
	}, nil
}

func (e *StreamedContainerExecutor) ExecuteCommand(cmd string) (Result, error) {

	var (
		outBuf = new(bytes.Buffer)
		errBuf = new(bytes.Buffer)
		delim  = e.delim
		err    error
	)

	fullCmd := fmt.Sprintf("%s; echo %s; echo %s 1>&2\n", cmd, delim, delim)

	if _, err = io.Copy(e.stdin, strings.NewReader(fullCmd)); err != nil {
		return Result{}, err
	}

	var wg concurrency.AsyncGroup
	wg.Do(func() { _, _ = io.Copy(outBuf, e.stdout) })
	wg.Do(func() { _, _ = io.Copy(errBuf, e.stderr) })
	wg.Wait()

	res := Result{
		Cmd:    cmd,
		Stdout: outBuf.String(),
		Stderr: errBuf.String(),
	}
	return res, nil
}
