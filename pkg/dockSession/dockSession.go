package dockSession

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/command"
	"dockgen/pkg/dockCopy"
	"dockgen/pkg/playback"
	"dockgen/pkg/recorder"
	"dockgen/pkg/runner"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"os"
)

type Session struct {
	historyPath string
	imageTag    string
	dockerfile  string
	client      *client.Client
	volume      volume.Volume

	playbackCmd []string
}

type SessionOptions struct {
	HistoryPath string
	ImageTag    string
	Dockerfile  string
}

func NewSession(cli *client.Client, opt SessionOptions) *Session {
	return &Session{
		historyPath: opt.HistoryPath,
		imageTag:    opt.ImageTag,
		dockerfile:  opt.Dockerfile,
		client:      cli,
		playbackCmd: nil,
	}
}
func (ds *Session) AddCmd(cmd []string) {
	ds.playbackCmd = append(ds.playbackCmd, cmd...)
}

func (ds *Session) createContainer(tty bool, raw bool) (*genContainer.Container, error) {
	return genContainer.NewContainer(ds.client, genContainer.CreateBuildAndContainerOptions{
		Tag:        ds.imageTag,
		Cmd:        "/bin/sh",
		Tty:        tty,
		Mounts:     map[string]string{ds.historyPath: "/root/.bash_history"},
		Env:        nil,
		DockerFile: ds.dockerfile,
		Raw:        raw,
	})
}

func (ds *Session) Run() ([]*command.Result, error) {
	_, err := os.Create(ds.historyPath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(ds.historyPath)

	interactiveContainer, err := ds.createContainer(true, true)
	if err != nil {
		return nil, err
	}

	go dockCopy.Copy(interactiveContainer.AttachContainer.Conn, os.Stdin)
	dockCopy.Copy(os.Stdout, interactiveContainer.AttachContainer.Reader)
	interactiveContainer.Close()

	rec := recorder.NewContainer(interactiveContainer, ds.historyPath)

	replayContainer, err := ds.createContainer(false, false)
	defer replayContainer.Close()
	if err != nil {
		return nil, err
	}

	pb, err := playback.NewContainer(replayContainer)
	if err != nil {
		return nil, err
	}

	r := runner.NewDockRunner(rec, pb, ds.playbackCmd)

	return r.Run()
}
