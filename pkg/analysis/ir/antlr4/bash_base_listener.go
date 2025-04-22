// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bash

import "github.com/antlr4-go/antlr/v4"

// BaseBashListener is a complete listener for a parse tree produced by BashParser.
type BaseBashListener struct{}

var _ BashListener = &BaseBashListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBashListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBashListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBashListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBashListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCommandLine is called when production commandLine is entered.
func (s *BaseBashListener) EnterCommandLine(ctx *CommandLineContext) {}

// ExitCommandLine is called when production commandLine is exited.
func (s *BaseBashListener) ExitCommandLine(ctx *CommandLineContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BaseBashListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BaseBashListener) ExitPipeline(ctx *PipelineContext) {}

// EnterOptionWithArg is called when production optionWithArg is entered.
func (s *BaseBashListener) EnterOptionWithArg(ctx *OptionWithArgContext) {}

// ExitOptionWithArg is called when production optionWithArg is exited.
func (s *BaseBashListener) ExitOptionWithArg(ctx *OptionWithArgContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseBashListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseBashListener) ExitCommand(ctx *CommandContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseBashListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseBashListener) ExitProg(ctx *ProgContext) {}

// EnterOption is called when production option is entered.
func (s *BaseBashListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseBashListener) ExitOption(ctx *OptionContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseBashListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseBashListener) ExitArg(ctx *ArgContext) {}

// EnterRedir is called when production redir is entered.
func (s *BaseBashListener) EnterRedir(ctx *RedirContext) {}

// ExitRedir is called when production redir is exited.
func (s *BaseBashListener) ExitRedir(ctx *RedirContext) {}

// EnterLogicalOp is called when production logicalOp is entered.
func (s *BaseBashListener) EnterLogicalOp(ctx *LogicalOpContext) {}

// ExitLogicalOp is called when production logicalOp is exited.
func (s *BaseBashListener) ExitLogicalOp(ctx *LogicalOpContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseBashListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseBashListener) ExitAssign(ctx *AssignContext) {}
