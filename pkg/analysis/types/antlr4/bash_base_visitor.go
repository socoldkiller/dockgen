// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bash

import "github.com/antlr4-go/antlr/v4"

type BaseBashVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBashVisitor) VisitCommandLine(ctx *CommandLineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitPipeline(ctx *PipelineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitOptionWithArg(ctx *OptionWithArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitRedir(ctx *RedirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitLogicalOp(ctx *LogicalOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBashVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}
