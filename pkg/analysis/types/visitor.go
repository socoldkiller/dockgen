package types

import (
	parser "dockgen/pkg/analysis/types/antlr4"
	"github.com/antlr4-go/antlr/v4"
)

func DefaultParseValue(val antlr.ParseTree) string {
	if val == nil {
		return ""
	}
	return val.GetText()
}

type IRVisitor struct {
	parser.BashVisitor
	defaultIR *AntlrCommandIR
}

func NewIRVisitor() *IRVisitor {
	return &IRVisitor{
		defaultIR: &AntlrCommandIR{},
	}
}

func (v *IRVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *IRVisitor) VisitCommandLine(ctx *parser.CommandLineContext) interface{} {
	if ctx.Pipeline() != nil {
		return v.Visit(ctx.Pipeline())
	}
	return &AntlrCommandIR{}
}

func (v *IRVisitor) VisitPipeline(ctx *parser.PipelineContext) interface{} {
	commands := ctx.AllCommand()
	ir := v.defaultIR
	if len(commands) == 1 {
		cmd := v.Visit(commands[0]).(*AntlrCommandIR)
		ir = cmd
		return ir
	}

	for _, cmdCtx := range commands {
		cmd := v.Visit(cmdCtx).(*AntlrCommandIR)
		ir.pipeCommand = append(v.defaultIR.pipeCommand, cmd)
	}

	return ir
}

func (v *IRVisitor) VisitCommand(ctx *parser.CommandContext) interface{} {
	var (
		env        Env
		Program    = ctx.Prog().GetText()
		options    = ctx.AllOptionWithArg()
		assign     = ctx.Assign()
		cmdOptions = make(map[string]string)
	)

	for _, opt := range options {
		key := opt.Option().GetText()
		value := DefaultParseValue(opt.Arg())
		cmdOptions[key] = value
	}

	var args []string
	for _, arg := range ctx.AllArg() {
		args = append(args, arg.GetText())
	}

	ioPut := make(map[string]*string)
	if len(ctx.AllRedir()) != 0 {
		last := ctx.AllRedir()[len(ctx.AllRedir())-1]
		var text string
		text = last.WORD().GetText()
		if last.LT() != nil {
			ioPut["input"] = &text
		}

		if last.GT() != nil || last.DGT() != nil {
			ioPut["output"] = &text
		}
	}

	ir := &AntlrCommandIR{
		program: Program,
		options: cmdOptions,
		args:    args,
		env:     nil,
		input:   ioPut["input"],
		output:  ioPut["output"],
	}

	if assign == nil || assign.GetChildCount() != 3 {
		return ir
	}

	env.EnvVariable = assign.WORD(0).GetText()
	var value Variable
	switch val := assign.GetChild(2).(type) {
	case *antlr.TerminalNodeImpl:
		value.Value = val.GetText()
	case *parser.VariableContext:
		value.Name = assign.Variable().GetText()
		value.Value = ""
	default:
	}

	env.EnvValue = value
	ir.env = &env
	return ir

}
