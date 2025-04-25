package types

import (
	parser "dockgen/pkg/analysis/types/antlr4"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type AntlrCommandIR struct {
	program     string
	pipeCommand []*AntlrCommandIR
	options     map[string]string
	args        []string
	input       *string
	output      *string
	redirect    *string
	env         *Env
}
type BashCommandIR struct {
	*AntlrCommandIR
	id     int
	cmd    string
	stdout *string
	stderr *string
}

func NewBashCommandIR(id int, cmd string, stdout, stderr *string) (*BashCommandIR, error) {
	input := antlr.NewInputStream(cmd)
	lexer := parser.NewBashLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewBashParser(tokens)
	tree := p.CommandLine()
	builder := NewIRVisitor()
	antlrIr, ok := builder.Visit(tree).(*AntlrCommandIR)
	if !ok {
		return nil, fmt.Errorf("generate bash IR error: %v", ok)
	}

	ir := &BashCommandIR{
		id:             id,
		cmd:            cmd,
		stdout:         stdout,
		stderr:         stderr,
		AntlrCommandIR: antlrIr,
	}

	return ir, nil
}

func (ir BashCommandIR) ID() int {
	return ir.id
}

func (ir BashCommandIR) Stdout() string {
	return *ir.stdout
}

func (ir BashCommandIR) Stderr() string {
	return *ir.stderr
}

func (ir BashCommandIR) Program() string {
	return ir.program
}

func (ir BashCommandIR) PipeCommand() []IR {
	var pipe []IR
	for _, p := range ir.pipeCommand {
		ir := &BashCommandIR{AntlrCommandIR: p}
		pipe = append(pipe, ir)
	}
	return pipe
}

func (ir BashCommandIR) Options() map[string]string {
	return ir.options
}

func (ir BashCommandIR) Args() []string {
	return ir.args
}

func (ir BashCommandIR) Input() *string {
	return ir.input
}

func (ir BashCommandIR) Output() *string {
	return ir.output
}

func (ir BashCommandIR) Redirect() *string {
	return ir.redirect
}

func (ir BashCommandIR) Env() *Env {
	return ir.env
}

func (ir BashCommandIR) Cmd() string {
	return ir.cmd
}
