// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Bash

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BashParser struct {
	*antlr.BaseParser
}

var BashParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bashParserInit() {
	staticData := &BashParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'|'", "'<'", "'>'", "'>>'", "'<>'", "'&&'", "'||'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT", "ANDAND", "OROR", "SEMI",
		"OPTION", "WORD", "WS",
	}
	staticData.RuleNames = []string{
		"commandLine", "pipeline", "optionWithArg", "command", "prog", "option",
		"arg", "redir", "logicalOp", "assign",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 86, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9, 1, 1, 2,
		1, 2, 3, 2, 34, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 40, 8, 3, 10, 3, 12,
		3, 43, 9, 3, 1, 3, 1, 3, 5, 3, 47, 8, 3, 10, 3, 12, 3, 50, 9, 3, 1, 3,
		5, 3, 53, 8, 3, 10, 3, 12, 3, 56, 9, 3, 1, 3, 5, 3, 59, 8, 3, 10, 3, 12,
		3, 62, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3, 67, 8, 3, 3, 3, 69, 8, 3, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 2, 1, 0,
		3, 6, 1, 0, 7, 9, 83, 0, 20, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 31, 1, 0,
		0, 0, 6, 68, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 74,
		1, 0, 0, 0, 14, 76, 1, 0, 0, 0, 16, 79, 1, 0, 0, 0, 18, 81, 1, 0, 0, 0,
		20, 21, 3, 2, 1, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 28, 3, 6,
		3, 0, 24, 25, 5, 2, 0, 0, 25, 27, 3, 6, 3, 0, 26, 24, 1, 0, 0, 0, 27, 30,
		1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 3, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 31, 33, 3, 10, 5, 0, 32, 34, 3, 12, 6, 0, 33, 32, 1,
		0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 5, 1, 0, 0, 0, 35, 36, 3, 8, 4, 0, 36,
		37, 3, 18, 9, 0, 37, 69, 1, 0, 0, 0, 38, 40, 3, 14, 7, 0, 39, 38, 1, 0,
		0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 44,
		1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 48, 3, 8, 4, 0, 45, 47, 3, 4, 2, 0,
		46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1,
		0, 0, 0, 49, 54, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 53, 3, 12, 6, 0, 52,
		51, 1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0,
		0, 55, 60, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 59, 3, 14, 7, 0, 58, 57,
		1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0,
		61, 66, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 3, 16, 8, 0, 64, 65, 3,
		6, 3, 0, 65, 67, 1, 0, 0, 0, 66, 63, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67,
		69, 1, 0, 0, 0, 68, 35, 1, 0, 0, 0, 68, 41, 1, 0, 0, 0, 69, 7, 1, 0, 0,
		0, 70, 71, 5, 11, 0, 0, 71, 9, 1, 0, 0, 0, 72, 73, 5, 10, 0, 0, 73, 11,
		1, 0, 0, 0, 74, 75, 5, 11, 0, 0, 75, 13, 1, 0, 0, 0, 76, 77, 7, 0, 0, 0,
		77, 78, 5, 11, 0, 0, 78, 15, 1, 0, 0, 0, 79, 80, 7, 1, 0, 0, 80, 17, 1,
		0, 0, 0, 81, 82, 5, 11, 0, 0, 82, 83, 5, 1, 0, 0, 83, 84, 5, 11, 0, 0,
		84, 19, 1, 0, 0, 0, 8, 28, 33, 41, 48, 54, 60, 66, 68,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// BashParserInit initializes any static state used to implement BashParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBashParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BashParserInit() {
	staticData := &BashParserStaticData
	staticData.once.Do(bashParserInit)
}

// NewBashParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBashParser(input antlr.TokenStream) *BashParser {
	BashParserInit()
	this := new(BashParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BashParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Bash.g4"

	return this
}

// BashParser tokens.
const (
	BashParserEOF    = antlr.TokenEOF
	BashParserASSIGN = 1
	BashParserPIPE   = 2
	BashParserLT     = 3
	BashParserGT     = 4
	BashParserDGT    = 5
	BashParserLTGT   = 6
	BashParserANDAND = 7
	BashParserOROR   = 8
	BashParserSEMI   = 9
	BashParserOPTION = 10
	BashParserWORD   = 11
	BashParserWS     = 12
)

// BashParser rules.
const (
	BashParserRULE_commandLine   = 0
	BashParserRULE_pipeline      = 1
	BashParserRULE_optionWithArg = 2
	BashParserRULE_command       = 3
	BashParserRULE_prog          = 4
	BashParserRULE_option        = 5
	BashParserRULE_arg           = 6
	BashParserRULE_redir         = 7
	BashParserRULE_logicalOp     = 8
	BashParserRULE_assign        = 9
)

// ICommandLineContext is an interface to support dynamic dispatch.
type ICommandLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Pipeline() IPipelineContext
	EOF() antlr.TerminalNode

	// IsCommandLineContext differentiates from other interfaces.
	IsCommandLineContext()
}

type CommandLineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandLineContext() *CommandLineContext {
	var p = new(CommandLineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_commandLine
	return p
}

func InitEmptyCommandLineContext(p *CommandLineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_commandLine
}

func (*CommandLineContext) IsCommandLineContext() {}

func NewCommandLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandLineContext {
	var p = new(CommandLineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_commandLine

	return p
}

func (s *CommandLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandLineContext) Pipeline() IPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *CommandLineContext) EOF() antlr.TerminalNode {
	return s.GetToken(BashParserEOF, 0)
}

func (s *CommandLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterCommandLine(s)
	}
}

func (s *CommandLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitCommandLine(s)
	}
}

func (s *CommandLineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitCommandLine(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) CommandLine() (localctx ICommandLineContext) {
	localctx = NewCommandLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BashParserRULE_commandLine)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(20)
		p.Pipeline()
	}
	{
		p.SetState(21)
		p.Match(BashParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCommand() []ICommandContext
	Command(i int) ICommandContext
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_pipeline
	return p
}

func InitEmptyPipelineContext(p *PipelineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_pipeline
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_pipeline

	return p
}

func (s *PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *PipelineContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(BashParserPIPE)
}

func (s *PipelineContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(BashParserPIPE, i)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (s *PipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BashParserRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)
		p.Command()
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BashParserPIPE {
		{
			p.SetState(24)
			p.Match(BashParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.Command()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionWithArgContext is an interface to support dynamic dispatch.
type IOptionWithArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Option() IOptionContext
	Arg() IArgContext

	// IsOptionWithArgContext differentiates from other interfaces.
	IsOptionWithArgContext()
}

type OptionWithArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionWithArgContext() *OptionWithArgContext {
	var p = new(OptionWithArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_optionWithArg
	return p
}

func InitEmptyOptionWithArgContext(p *OptionWithArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_optionWithArg
}

func (*OptionWithArgContext) IsOptionWithArgContext() {}

func NewOptionWithArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionWithArgContext {
	var p = new(OptionWithArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_optionWithArg

	return p
}

func (s *OptionWithArgContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionWithArgContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *OptionWithArgContext) Arg() IArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *OptionWithArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionWithArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionWithArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterOptionWithArg(s)
	}
}

func (s *OptionWithArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitOptionWithArg(s)
	}
}

func (s *OptionWithArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitOptionWithArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) OptionWithArg() (localctx IOptionWithArgContext) {
	localctx = NewOptionWithArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BashParserRULE_optionWithArg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Option()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(32)
			p.Arg()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prog() IProgContext
	Assign() IAssignContext
	AllRedir() []IRedirContext
	Redir(i int) IRedirContext
	AllOptionWithArg() []IOptionWithArgContext
	OptionWithArg(i int) IOptionWithArgContext
	AllArg() []IArgContext
	Arg(i int) IArgContext
	LogicalOp() ILogicalOpContext
	Command() ICommandContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Prog() IProgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgContext)
}

func (s *CommandContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *CommandContext) AllRedir() []IRedirContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRedirContext); ok {
			len++
		}
	}

	tst := make([]IRedirContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRedirContext); ok {
			tst[i] = t.(IRedirContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Redir(i int) IRedirContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRedirContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRedirContext)
}

func (s *CommandContext) AllOptionWithArg() []IOptionWithArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionWithArgContext); ok {
			len++
		}
	}

	tst := make([]IOptionWithArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionWithArgContext); ok {
			tst[i] = t.(IOptionWithArgContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) OptionWithArg(i int) IOptionWithArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionWithArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionWithArgContext)
}

func (s *CommandContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *CommandContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *CommandContext) LogicalOp() ILogicalOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOpContext)
}

func (s *CommandContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BashParserRULE_command)
	var _la int

	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Prog()
		}
		{
			p.SetState(36)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0 {
			{
				p.SetState(38)
				p.Redir()
			}

			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(44)
			p.Prog()
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BashParserOPTION {
			{
				p.SetState(45)
				p.OptionWithArg()
			}

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BashParserWORD {
			{
				p.SetState(51)
				p.Arg()
			}

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0 {
			{
				p.SetState(57)
				p.Redir()
			}

			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0 {
			{
				p.SetState(63)
				p.LogicalOp()
			}
			{
				p.SetState(64)
				p.Command()
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) WORD() antlr.TerminalNode {
	return s.GetToken(BashParserWORD, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BashParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) OPTION() antlr.TerminalNode {
	return s.GetToken(BashParserOPTION, 0)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitOption(s)
	}
}

func (s *OptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BashParserRULE_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(BashParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_arg
	return p
}

func InitEmptyArgContext(p *ArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_arg
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) WORD() antlr.TerminalNode {
	return s.GetToken(BashParserWORD, 0)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitArg(s)
	}
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BashParserRULE_arg)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRedirContext is an interface to support dynamic dispatch.
type IRedirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	DGT() antlr.TerminalNode
	LTGT() antlr.TerminalNode

	// IsRedirContext differentiates from other interfaces.
	IsRedirContext()
}

type RedirContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRedirContext() *RedirContext {
	var p = new(RedirContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_redir
	return p
}

func InitEmptyRedirContext(p *RedirContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_redir
}

func (*RedirContext) IsRedirContext() {}

func NewRedirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RedirContext {
	var p = new(RedirContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_redir

	return p
}

func (s *RedirContext) GetParser() antlr.Parser { return s.parser }

func (s *RedirContext) WORD() antlr.TerminalNode {
	return s.GetToken(BashParserWORD, 0)
}

func (s *RedirContext) LT() antlr.TerminalNode {
	return s.GetToken(BashParserLT, 0)
}

func (s *RedirContext) GT() antlr.TerminalNode {
	return s.GetToken(BashParserGT, 0)
}

func (s *RedirContext) DGT() antlr.TerminalNode {
	return s.GetToken(BashParserDGT, 0)
}

func (s *RedirContext) LTGT() antlr.TerminalNode {
	return s.GetToken(BashParserLTGT, 0)
}

func (s *RedirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RedirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RedirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterRedir(s)
	}
}

func (s *RedirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitRedir(s)
	}
}

func (s *RedirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitRedir(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Redir() (localctx IRedirContext) {
	localctx = NewRedirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BashParserRULE_redir)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&120) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(77)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogicalOpContext is an interface to support dynamic dispatch.
type ILogicalOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ANDAND() antlr.TerminalNode
	OROR() antlr.TerminalNode
	SEMI() antlr.TerminalNode

	// IsLogicalOpContext differentiates from other interfaces.
	IsLogicalOpContext()
}

type LogicalOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOpContext() *LogicalOpContext {
	var p = new(LogicalOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_logicalOp
	return p
}

func InitEmptyLogicalOpContext(p *LogicalOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_logicalOp
}

func (*LogicalOpContext) IsLogicalOpContext() {}

func NewLogicalOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOpContext {
	var p = new(LogicalOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_logicalOp

	return p
}

func (s *LogicalOpContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOpContext) ANDAND() antlr.TerminalNode {
	return s.GetToken(BashParserANDAND, 0)
}

func (s *LogicalOpContext) OROR() antlr.TerminalNode {
	return s.GetToken(BashParserOROR, 0)
}

func (s *LogicalOpContext) SEMI() antlr.TerminalNode {
	return s.GetToken(BashParserSEMI, 0)
}

func (s *LogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterLogicalOp(s)
	}
}

func (s *LogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitLogicalOp(s)
	}
}

func (s *LogicalOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitLogicalOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) LogicalOp() (localctx ILogicalOpContext) {
	localctx = NewLogicalOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BashParserRULE_logicalOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllWORD() []antlr.TerminalNode
	WORD(i int) antlr.TerminalNode
	ASSIGN() antlr.TerminalNode

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(BashParserWORD)
}

func (s *AssignContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(BashParserWORD, i)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BashParserASSIGN, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BashParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(BashParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
