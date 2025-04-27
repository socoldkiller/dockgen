package runner

import (
	"dockgen/pkg/command"
	"dockgen/pkg/playback"
	"dockgen/pkg/recorder"
)

type DockRunner struct {
	recorder recorder.Recorder
	playback playback.PlayBack

	infoCmd []string
}

func NewDockRunner(recorder recorder.Recorder, playback playback.PlayBack, infoCmd []string) *DockRunner {
	return &DockRunner{
		recorder: recorder,
		playback: playback,
		infoCmd:  infoCmd,
	}
}

func (r DockRunner) Run() ([]*command.Result, error) {
	var (
		err     error
		cmdList []string
	)

	if cmdList, err = r.recorder.Record(); err != nil {
		return nil, err
	}

	cmdList = append(cmdList, r.infoCmd...)
	return r.playback.PlayBack(cmdList)
}
