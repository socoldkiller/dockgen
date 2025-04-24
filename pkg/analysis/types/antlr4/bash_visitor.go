// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bash

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BashParser.
type BashVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BashParser#commandLine.
	VisitCommandLine(ctx *CommandLineContext) interface{}

	// Visit a parse tree produced by BashParser#pipeline.
	VisitPipeline(ctx *PipelineContext) interface{}

	// Visit a parse tree produced by BashParser#optionWithArg.
	VisitOptionWithArg(ctx *OptionWithArgContext) interface{}

	// Visit a parse tree produced by BashParser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by BashParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by BashParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by BashParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by BashParser#redir.
	VisitRedir(ctx *RedirContext) interface{}

	// Visit a parse tree produced by BashParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by BashParser#logicalOp.
	VisitLogicalOp(ctx *LogicalOpContext) interface{}

	// Visit a parse tree produced by BashParser#assign.
	VisitAssign(ctx *AssignContext) interface{}
}
