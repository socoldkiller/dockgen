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
		"", "'='", "'|'", "'<'", "'>'", "'>>'", "'<>'", "'&&'", "'||'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT", "ANDAND", "OROR", "SEMI",
		"OPTION", "WORD", "WS",
	}
	staticData.RuleNames = []string{
		"ASSIGN", "PIPE", "LT", "GT", "DGT", "LTGT", "ANDAND", "OROR", "SEMI",
		"DASH", "OPTION", "WORD", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 73, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10, 53, 8, 10, 11, 10,
		12, 10, 54, 1, 10, 4, 10, 58, 8, 10, 11, 10, 12, 10, 59, 1, 11, 4, 11,
		63, 8, 11, 11, 11, 12, 11, 64, 1, 12, 4, 12, 68, 8, 12, 11, 12, 12, 12,
		69, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 0, 21, 10, 23, 11, 25, 12, 1, 0, 3, 11, 0, 9, 10, 13,
		13, 32, 32, 34, 34, 36, 36, 38, 41, 45, 45, 59, 62, 92, 92, 96, 96, 124,
		124, 10, 0, 9, 10, 13, 13, 32, 32, 36, 36, 38, 38, 40, 41, 59, 62, 92,
		92, 96, 96, 124, 124, 3, 0, 9, 10, 13, 13, 32, 32, 75, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0,
		0, 0, 3, 29, 1, 0, 0, 0, 5, 31, 1, 0, 0, 0, 7, 33, 1, 0, 0, 0, 9, 35, 1,
		0, 0, 0, 11, 38, 1, 0, 0, 0, 13, 41, 1, 0, 0, 0, 15, 44, 1, 0, 0, 0, 17,
		47, 1, 0, 0, 0, 19, 49, 1, 0, 0, 0, 21, 52, 1, 0, 0, 0, 23, 62, 1, 0, 0,
		0, 25, 67, 1, 0, 0, 0, 27, 28, 5, 61, 0, 0, 28, 2, 1, 0, 0, 0, 29, 30,
		5, 124, 0, 0, 30, 4, 1, 0, 0, 0, 31, 32, 5, 60, 0, 0, 32, 6, 1, 0, 0, 0,
		33, 34, 5, 62, 0, 0, 34, 8, 1, 0, 0, 0, 35, 36, 5, 62, 0, 0, 36, 37, 5,
		62, 0, 0, 37, 10, 1, 0, 0, 0, 38, 39, 5, 60, 0, 0, 39, 40, 5, 62, 0, 0,
		40, 12, 1, 0, 0, 0, 41, 42, 5, 38, 0, 0, 42, 43, 5, 38, 0, 0, 43, 14, 1,
		0, 0, 0, 44, 45, 5, 124, 0, 0, 45, 46, 5, 124, 0, 0, 46, 16, 1, 0, 0, 0,
		47, 48, 5, 59, 0, 0, 48, 18, 1, 0, 0, 0, 49, 50, 5, 45, 0, 0, 50, 20, 1,
		0, 0, 0, 51, 53, 3, 19, 9, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 58, 8, 0, 0,
		0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60,
		1, 0, 0, 0, 60, 22, 1, 0, 0, 0, 61, 63, 8, 1, 0, 0, 62, 61, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 24, 1,
		0, 0, 0, 66, 68, 7, 2, 0, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69,
		67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 6, 12,
		0, 0, 72, 26, 1, 0, 0, 0, 5, 0, 54, 59, 64, 69, 1, 6, 0, 0,
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
	BashLexerASSIGN = 1
	BashLexerPIPE   = 2
	BashLexerLT     = 3
	BashLexerGT     = 4
	BashLexerDGT    = 5
	BashLexerLTGT   = 6
	BashLexerANDAND = 7
	BashLexerOROR   = 8
	BashLexerSEMI   = 9
	BashLexerOPTION = 10
	BashLexerWORD   = 11
	BashLexerWS     = 12
)
