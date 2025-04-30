// Code generated from Bash.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BashLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BashLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bashlexerLexerInit() {
	staticData := &BashLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'$'", "'('", "')'", "'='", "'|'", "'<'", "'>'", "'>>'", "'<>'",
		"'&&'", "'||'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT", "ANDAND",
		"OROR", "SEMI", "OPTION", "WORD", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT",
		"ANDAND", "OROR", "SEMI", "DASH", "OPTION", "WORD", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 85, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 4, 13, 65, 8, 13,
		11, 13, 12, 13, 66, 1, 13, 4, 13, 70, 8, 13, 11, 13, 12, 13, 71, 1, 14,
		4, 14, 75, 8, 14, 11, 14, 12, 14, 76, 1, 15, 4, 15, 80, 8, 15, 11, 15,
		12, 15, 81, 1, 15, 1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0, 27, 13, 29, 14, 31,
		15, 1, 0, 3, 10, 0, 9, 10, 13, 13, 32, 32, 34, 34, 36, 36, 38, 41, 59,
		62, 92, 92, 96, 96, 124, 124, 9, 0, 9, 10, 13, 13, 32, 32, 38, 38, 40,
		41, 59, 62, 92, 92, 96, 96, 124, 124, 3, 0, 9, 10, 13, 13, 32, 32, 87,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 1, 33, 1,
		0, 0, 0, 3, 35, 1, 0, 0, 0, 5, 37, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 41,
		1, 0, 0, 0, 11, 43, 1, 0, 0, 0, 13, 45, 1, 0, 0, 0, 15, 47, 1, 0, 0, 0,
		17, 50, 1, 0, 0, 0, 19, 53, 1, 0, 0, 0, 21, 56, 1, 0, 0, 0, 23, 59, 1,
		0, 0, 0, 25, 61, 1, 0, 0, 0, 27, 64, 1, 0, 0, 0, 29, 74, 1, 0, 0, 0, 31,
		79, 1, 0, 0, 0, 33, 34, 5, 36, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 40,
		0, 0, 36, 4, 1, 0, 0, 0, 37, 38, 5, 41, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40,
		5, 61, 0, 0, 40, 8, 1, 0, 0, 0, 41, 42, 5, 124, 0, 0, 42, 10, 1, 0, 0,
		0, 43, 44, 5, 60, 0, 0, 44, 12, 1, 0, 0, 0, 45, 46, 5, 62, 0, 0, 46, 14,
		1, 0, 0, 0, 47, 48, 5, 62, 0, 0, 48, 49, 5, 62, 0, 0, 49, 16, 1, 0, 0,
		0, 50, 51, 5, 60, 0, 0, 51, 52, 5, 62, 0, 0, 52, 18, 1, 0, 0, 0, 53, 54,
		5, 38, 0, 0, 54, 55, 5, 38, 0, 0, 55, 20, 1, 0, 0, 0, 56, 57, 5, 124, 0,
		0, 57, 58, 5, 124, 0, 0, 58, 22, 1, 0, 0, 0, 59, 60, 5, 59, 0, 0, 60, 24,
		1, 0, 0, 0, 61, 62, 5, 45, 0, 0, 62, 26, 1, 0, 0, 0, 63, 65, 3, 25, 12,
		0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 70, 8, 0, 0, 0, 69, 68, 1, 0, 0, 0,
		70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 28, 1,
		0, 0, 0, 73, 75, 8, 1, 0, 0, 74, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 30, 1, 0, 0, 0, 78, 80, 7, 2, 0,
		0, 79, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82,
		1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 6, 15, 0, 0, 84, 32, 1, 0, 0, 0,
		5, 0, 66, 71, 76, 81, 1, 6, 0, 0,
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

// BashLexerInit initializes any static state used to implement BashLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBashLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BashLexerInit() {
	staticData := &BashLexerLexerStaticData
	staticData.once.Do(bashlexerLexerInit)
}

// NewBashLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBashLexer(input antlr.CharStream) *BashLexer {
	BashLexerInit()
	l := new(BashLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BashLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Bash.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BashLexer tokens.
const (
	BashLexerT__0   = 1
	BashLexerT__1   = 2
	BashLexerT__2   = 3
	BashLexerASSIGN = 4
	BashLexerPIPE   = 5
	BashLexerLT     = 6
	BashLexerGT     = 7
	BashLexerDGT    = 8
	BashLexerLTGT   = 9
	BashLexerANDAND = 10
	BashLexerOROR   = 11
	BashLexerSEMI   = 12
	BashLexerOPTION = 13
	BashLexerWORD   = 14
	BashLexerWS     = 15
)
