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
}

func NewCommandIR(cmd *command.Result) *CommandIR {
	res := strings.SplitN(cmd.Cmd, " ", 2)
	return &CommandIR{
		Raw:    res[0],
		Stdout: cmd.Stdout,
		Stderr: cmd.Stderr,
		Args:   res[1:],
		Tty:    false,
	}
}

type IREdge struct {
	from   int
	to     int
	Tags   []string
	Weight int
}

func AddIREdge(from, to int, tags []string) *IREdge {
	return &IREdge{
		from: from,
		to:   to,
		Tags: tags,
	}
}
