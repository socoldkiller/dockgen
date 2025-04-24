package ir

import (
	"dockgen/pkg/util"
)

type Variable struct {
	Name  string
	Value string
}

type Env struct {
	EnvVariable string
	EnvValue    Variable
}

type BashCommandIR struct {
	id          int
	cmd         string
	stdout      *string
	stderr      *string
	program     string
	pipeCommand []*BashCommandIR
	options     map[string]string
	args        []string
	input       *string
	output      *string
	redir       string
	env         *Env
}

func (ir BashCommandIR) ToMap() map[string]interface{} {
	data := map[string]interface{}{
		"cmd":         ir.cmd,
		"id":          ir.id,
		"stdout":      ir.stdout,
		"stderr":      ir.stderr,
		"program":     ir.program,
		"pipeCommand": ir.pipeCommand,
		"options":     ir.options,
		"args":        ir.args,
		"input":       ir.input,
		"output":      ir.output,
		"redir":       ir.redir,
		"env":         ir.env,
	}

	return data
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
	pipe, _ := util.CastSlice[*BashCommandIR, IR](ir.pipeCommand)
	return pipe
}

func (ir BashCommandIR) Options() map[string]string {
	return ir.options
}

func (ir BashCommandIR) Args() []string {
	return ir.args
}

func (ir BashCommandIR) Input() string {
	return *ir.input
}

func (ir BashCommandIR) Output() string {
	return *ir.output
}

func (ir BashCommandIR) Redir() string {
	return ir.redir
}

func (ir BashCommandIR) Env() Env {
	return *ir.env
}

type IR interface {
	ID() int
	Stdout() string
	Stderr() string
	Program() string
	PipeCommand() []IR
	Options() map[string]string
	Args() []string
	Input() string
	Output() string
	Redir() string
	Env() Env
}
