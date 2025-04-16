package runner

import (
	"dockgen/pkg/command"
	"dockgen/pkg/playback"
	"dockgen/pkg/recorder"
)

type DockRunner struct {
	recorder recorder.Recorder
	playback playback.PlayBack
}

func NewDockRunner(recorder recorder.Recorder, playback playback.PlayBack) *DockRunner {
	return &DockRunner{recorder: recorder, playback: playback}
}

func (r DockRunner) Run() ([]command.Result, error) {
	var (
		err     error
		cmdList []string
	)

	if cmdList, err = r.recorder.Record(); err != nil {
		return nil, err
	}

	return r.playback.PlayBack(cmdList)
}
