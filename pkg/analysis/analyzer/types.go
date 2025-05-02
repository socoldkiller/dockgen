package analyzer

import (
	"dockgen/pkg/analysis/types"
)

type Analyzer interface {
	Init(az ...Analyzer)
	Analyze(graph types.CFGraph) (ResultReader, error)
	Reset()
}

type ResultReader interface {
	Name() string // analyzer provider name
	Raw() []string
	Items() []Item
	QueryIDIR(int) (Item, bool)
	AnalyzeRaw() []string
}

type Item struct {
	OldIR types.IR
	NewIR types.IR
}
