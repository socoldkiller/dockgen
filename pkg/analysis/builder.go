package analysis

import (
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
)

type IRBuilder interface {
	Build([]*command.Result) ([]types.IR, error)
}

type BashIRBuilder struct {
}

func NewBashIRBuilder() *BashIRBuilder {
	return &BashIRBuilder{}
}

func (c BashIRBuilder) buildCommandIR(cmd string, stderr, stdout string, id int) (cmdIR *types.BashCommandIR, err error) {
	return types.NewBashCommandIR(id, cmd, &stdout, &stderr)
}

func (c BashIRBuilder) Build(cmdList []*command.Result) ([]types.IR, error) {
	var irList []types.IR
	for idx, cmd := range cmdList {
		cmdIR, err := c.buildCommandIR(cmd.Cmd, cmd.Stderr, cmd.Stdout, idx+1)
		if err != nil {
			return nil, err
		}
		irList = append(irList, cmdIR)
	}
	return irList, nil
}
