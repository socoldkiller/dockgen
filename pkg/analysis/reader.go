package analysis

import (
	"dockgen/pkg/command"
	"dockgen/pkg/playback"
)

type CmdReader interface {
	Read() ([]*command.Result, error)
}

type PlayBackCmdReader struct {
	playback playback.PlayBack
	cmdList  []string
}

func NewPlayBackCmdReader(playback playback.PlayBack, cmdList []string) *PlayBackCmdReader {
	return &PlayBackCmdReader{playback: playback, cmdList: cmdList}
}

func (p PlayBackCmdReader) Read() ([]*command.Result, error) {
	return p.playback.PlayBack(p.cmdList)

}
