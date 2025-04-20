package analysis

import (
	"dockgen/pkg/command"
	"strings"
)

type CommandIR struct {
	ID         int
	Raw        string   `json:"raw"`
	Stdout     string   `json:"stdout"`
	Stderr     string   `json:"stderr"`
	Executable string   `json:"executable"`
	Tty        bool     `json:"tty"`
	Args       []string `json:"args"`
	Cmd        string   `json:"cmd"`
}

func NewCommandIR(cmd *command.Result) *CommandIR {
	res := strings.SplitN(cmd.Cmd, " ", 2)
	return &CommandIR{
		Cmd:    cmd.Cmd,
		Raw:    res[0],
		Stdout: cmd.Stdout,
		Stderr: cmd.Stderr,
		Args:   res[1:],
		Tty:    false,
	}
}
