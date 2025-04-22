package ir

import (
	parser "dockgen/pkg/analysis/ir/antlr4"
	"github.com/antlr4-go/antlr/v4"
)

type Env struct {
	Var string
	Val string
}

type BashCommandIR struct {
	ID          int
	Stdout      string
	Stderr      string
	Program     string
	PipeCommand []*BashCommandIR
	Options     map[string]string
	Args        []string
	Input       *string
	Output      *string
	Redir       string
	env         Env
}

type IRBuilder struct {
	parser.BashVisitor
	IR *BashCommandIR
}

func NewIRBuilder() *IRBuilder {
	return &IRBuilder{
		IR: &BashCommandIR{},
	}
}

func (v *IRBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *IRBuilder) VisitCommandLine(ctx *parser.CommandLineContext) interface{} {
	return v.Visit(ctx.Pipeline())
}

func (v *IRBuilder) VisitPipeline(ctx *parser.PipelineContext) interface{} {
	commands := ctx.AllCommand()
	if len(commands) == 1 {
		cmd := v.Visit(commands[0]).(*BashCommandIR)
		v.IR = cmd
		return v.IR
	}

	for _, cmdCtx := range commands {
		cmd := v.Visit(cmdCtx).(*BashCommandIR)
		v.IR.PipeCommand = append(v.IR.PipeCommand, cmd)
	}
	return v.IR
}

func (v *IRBuilder) VisitCommand(ctx *parser.CommandContext) interface{} {
	Program := ctx.Prog().GetText()

	options := ctx.AllOptionWithArg()
	cmdOptions := make(map[string]string)
	for _, opt := range options {
		key := opt.Option().GetText()
		value := opt.Arg().GetText()
		cmdOptions[key] = value
	}

	var args []string
	for _, arg := range ctx.AllArg() {
		args = append(args, arg.GetText())
	}

	var env Env

	if ctx.Assign() != nil {
		assign := ctx.Assign()
		env.Var = assign.WORD(0).GetText()
		env.Val = assign.WORD(1).GetText()
	}

	return &BashCommandIR{
		Program: Program,
		Options: cmdOptions,
		Args:    args,
		env:     env,
	}
}
