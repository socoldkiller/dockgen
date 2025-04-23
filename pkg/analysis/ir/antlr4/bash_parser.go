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
		"", "'$'", "'('", "')'", "'='", "'|'", "'<'", "'>'", "'>>'", "'<>'",
		"'&&'", "'||'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT", "ANDAND",
		"OROR", "SEMI", "OPTION", "WORD", "WS",
	}
	staticData.RuleNames = []string{
		"commandLine", "pipeline", "optionWithArg", "command", "prog", "option",
		"arg", "redir", "variable", "logicalOp", "assign",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 109, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 27, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 32,
		8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 45, 8, 3, 10, 3, 12, 3, 48, 9, 3, 1, 3, 1, 3, 5, 3, 52,
		8, 3, 10, 3, 12, 3, 55, 9, 3, 1, 3, 5, 3, 58, 8, 3, 10, 3, 12, 3, 61, 9,
		3, 1, 3, 5, 3, 64, 8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 3, 1, 3, 1, 3, 3, 3,
		72, 8, 3, 3, 3, 74, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 82,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 5, 8, 89, 8, 8, 10, 8, 12, 8, 92, 9,
		8, 1, 8, 1, 8, 5, 8, 96, 8, 8, 10, 8, 12, 8, 99, 9, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 107, 8, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 0, 2, 1, 0, 6, 9, 1, 0, 10, 12, 110, 0, 26, 1,
		0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 36, 1, 0, 0, 0, 6, 73, 1, 0, 0, 0, 8, 75,
		1, 0, 0, 0, 10, 77, 1, 0, 0, 0, 12, 81, 1, 0, 0, 0, 14, 83, 1, 0, 0, 0,
		16, 86, 1, 0, 0, 0, 18, 100, 1, 0, 0, 0, 20, 102, 1, 0, 0, 0, 22, 23, 3,
		2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 27, 1, 0, 0, 0, 25, 27, 5, 0, 0, 1, 26,
		22, 1, 0, 0, 0, 26, 25, 1, 0, 0, 0, 27, 1, 1, 0, 0, 0, 28, 33, 3, 6, 3,
		0, 29, 30, 5, 5, 0, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1, 0, 0, 0, 32, 35,
		1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 3, 1, 0, 0, 0,
		35, 33, 1, 0, 0, 0, 36, 38, 3, 10, 5, 0, 37, 39, 3, 12, 6, 0, 38, 37, 1,
		0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 3, 8, 4, 0, 41,
		42, 3, 20, 10, 0, 42, 74, 1, 0, 0, 0, 43, 45, 3, 14, 7, 0, 44, 43, 1, 0,
		0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49,
		1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49, 53, 3, 8, 4, 0, 50, 52, 3, 4, 2, 0,
		51, 50, 1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1,
		0, 0, 0, 54, 59, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 58, 3, 12, 6, 0, 57,
		56, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0,
		0, 60, 65, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 64, 3, 14, 7, 0, 63, 62,
		1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 71, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 69, 3, 18, 9, 0, 69, 70, 3,
		6, 3, 0, 70, 72, 1, 0, 0, 0, 71, 68, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72,
		74, 1, 0, 0, 0, 73, 40, 1, 0, 0, 0, 73, 46, 1, 0, 0, 0, 74, 7, 1, 0, 0,
		0, 75, 76, 5, 14, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 5, 13, 0, 0, 78, 11,
		1, 0, 0, 0, 79, 82, 5, 14, 0, 0, 80, 82, 3, 16, 8, 0, 81, 79, 1, 0, 0,
		0, 81, 80, 1, 0, 0, 0, 82, 13, 1, 0, 0, 0, 83, 84, 7, 0, 0, 0, 84, 85,
		5, 14, 0, 0, 85, 15, 1, 0, 0, 0, 86, 90, 5, 1, 0, 0, 87, 89, 5, 2, 0, 0,
		88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1,
		0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 97, 5, 14, 0, 0, 94,
		96, 5, 3, 0, 0, 95, 94, 1, 0, 0, 0, 96, 99, 1, 0, 0, 0, 97, 95, 1, 0, 0,
		0, 97, 98, 1, 0, 0, 0, 98, 17, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 100, 101,
		7, 1, 0, 0, 101, 19, 1, 0, 0, 0, 102, 103, 5, 14, 0, 0, 103, 106, 5, 4,
		0, 0, 104, 107, 5, 14, 0, 0, 105, 107, 3, 16, 8, 0, 106, 104, 1, 0, 0,
		0, 106, 105, 1, 0, 0, 0, 107, 21, 1, 0, 0, 0, 13, 26, 33, 38, 46, 53, 59,
		65, 71, 73, 81, 90, 97, 106,
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
	BashParserT__0   = 1
	BashParserT__1   = 2
	BashParserT__2   = 3
	BashParserASSIGN = 4
	BashParserPIPE   = 5
	BashParserLT     = 6
	BashParserGT     = 7
	BashParserDGT    = 8
	BashParserLTGT   = 9
	BashParserANDAND = 10
	BashParserOROR   = 11
	BashParserSEMI   = 12
	BashParserOPTION = 13
	BashParserWORD   = 14
	BashParserWS     = 15
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
	BashParserRULE_variable      = 8
	BashParserRULE_logicalOp     = 9
	BashParserRULE_assign        = 10
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BashParserLT, BashParserGT, BashParserDGT, BashParserLTGT, BashParserWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Pipeline()
		}
		{
			p.SetState(23)
			p.Match(BashParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BashParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(BashParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(28)
		p.Command()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BashParserPIPE {
		{
			p.SetState(29)
			p.Match(BashParserPIPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Command()
		}

		p.SetState(35)
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
		p.SetState(36)
		p.Option()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(37)
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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Prog()
		}
		{
			p.SetState(41)
			p.Assign()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0 {
			{
				p.SetState(43)
				p.Redir()
			}

			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(49)
			p.Prog()
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BashParserOPTION {
			{
				p.SetState(50)
				p.OptionWithArg()
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BashParserT__0 || _la == BashParserWORD {
			{
				p.SetState(56)
				p.Arg()
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0 {
			{
				p.SetState(62)
				p.Redir()
			}

			p.SetState(67)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0 {
			{
				p.SetState(68)
				p.LogicalOp()
			}
			{
				p.SetState(69)
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
		p.SetState(75)
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
		p.SetState(77)
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
	Variable() IVariableContext

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

func (s *ArgContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BashParserWORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Match(BashParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BashParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Variable()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&960) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(84)
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

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BashParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BashParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) WORD() antlr.TerminalNode {
	return s.GetToken(BashParserWORD, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BashListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BashVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BashParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BashParserRULE_variable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(BashParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BashParserT__1 {
		{
			p.SetState(87)
			p.Match(BashParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BashParserT__2 {
		{
			p.SetState(94)
			p.Match(BashParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(99)
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
	p.EnterRule(localctx, 18, BashParserRULE_logicalOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
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
	Variable() IVariableContext

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

func (s *AssignContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
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
	p.EnterRule(localctx, 20, BashParserRULE_assign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(BashParserWORD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.Match(BashParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BashParserWORD:
		{
			p.SetState(104)
			p.Match(BashParserWORD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case BashParserT__0:
		{
			p.SetState(105)
			p.Variable()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
