// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bash

import "github.com/antlr4-go/antlr/v4"

// BashListener is a complete listener for a parse tree produced by BashParser.
type BashListener interface {
	antlr.ParseTreeListener

	// EnterCommandLine is called when entering the commandLine production.
	EnterCommandLine(c *CommandLineContext)

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterOptionWithArg is called when entering the optionWithArg production.
	EnterOptionWithArg(c *OptionWithArgContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterRedir is called when entering the redir production.
	EnterRedir(c *RedirContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterLogicalOp is called when entering the logicalOp production.
	EnterLogicalOp(c *LogicalOpContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// ExitCommandLine is called when exiting the commandLine production.
	ExitCommandLine(c *CommandLineContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitOptionWithArg is called when exiting the optionWithArg production.
	ExitOptionWithArg(c *OptionWithArgContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitRedir is called when exiting the redir production.
	ExitRedir(c *RedirContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitLogicalOp is called when exiting the logicalOp production.
	ExitLogicalOp(c *LogicalOpContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)
}
