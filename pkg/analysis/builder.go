package analysis

import (
	"dockgen/pkg/analysis/ir"
	parser "dockgen/pkg/analysis/ir/antlr4"
	"dockgen/pkg/command"
	"dockgen/pkg/log"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type IRBuilder interface {
	Build([]*command.Result) ([]*ir.BashCommandIR, error)
}

type CmdIRBuilder struct {
}

func NewCmdIRBuilder() *CmdIRBuilder {
	return &CmdIRBuilder{}
}

func (c CmdIRBuilder) buildCommandIR(cmd string, stderr, stdout string, id int) (cmdIR *ir.BashCommandIR, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Warnf("cmd: '%s' build command IR error %v", cmd, r)
			err = fmt.Errorf("panic while parsing command '%s': %v", cmd, r)
		}
	}()

	is := antlr.NewInputStream(cmd)
	lexer := parser.NewBashLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBashParser(tokens)

	tree := p.CommandLine()
	builder := ir.NewIRBuilder()
	cmdIR, ok := builder.Visit(tree).(*ir.BashCommandIR)
	if !ok {
		return nil, fmt.Errorf("failed to parse command: %s", cmd)
	}
	cmdIR.Stderr = stderr
	cmdIR.Stdout = stdout
	cmdIR.ID = id
	return cmdIR, nil
}

func (c CmdIRBuilder) Build(cmdList []*command.Result) ([]*ir.BashCommandIR, error) {
	var irList []*ir.BashCommandIR
	for idx, cmd := range cmdList {
		cmdIR, err := c.buildCommandIR(cmd.Cmd, cmd.Stderr, cmd.Stdout, idx+1)
		if err != nil {
			return nil, err
		}
		irList = append(irList, cmdIR)
	}
	return irList, nil
}
