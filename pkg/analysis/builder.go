package analysis

import (
	"dockgen/pkg/command"
)

type IRBuilder interface {
	Build([]*command.Result) ([]*CommandIR, error)
}

type CmdIRBuilder struct {
}

func NewCmdIRBuilder() *CmdIRBuilder {
	return &CmdIRBuilder{}
}

func (c CmdIRBuilder) Build(cmdList []*command.Result) ([]*CommandIR, error) {
	var irList []*CommandIR
	for idx, c := range cmdList {
		ir := NewCommandIR(c)
		ir.ID = idx + 1
		irList = append(irList, ir)
	}

	return irList, nil
}
