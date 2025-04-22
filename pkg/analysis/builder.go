package analysis

import (
	"dockgen/pkg/analysis/ir"
	parser "dockgen/pkg/analysis/ir/antlr4"
	"dockgen/pkg/command"
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

func (c CmdIRBuilder) Build(cmdList []*command.Result) ([]*ir.BashCommandIR, error) {
	var irList []*ir.BashCommandIR
	for idx, c := range cmdList {

		is := antlr.NewInputStream(c.Cmd)
		lexer := parser.NewBashLexer(is)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewBashParser(tokens)
		tree := p.CommandLine()
		builder := ir.NewIRBuilder()
		cmdIR, _ := builder.Visit(tree).(*ir.BashCommandIR)
		cmdIR.Stderr = c.Stderr
		cmdIR.Stdout = c.Stdout
		cmdIR.ID = idx + 1
		irList = append(irList, cmdIR)
	}

	return irList, nil
}
