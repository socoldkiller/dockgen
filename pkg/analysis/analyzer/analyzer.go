package analyzer

import (
	"dockgen/pkg/analysis/types"
)

type Analyzer interface {
	Analyze(graph types.CFGraph) (*Result, error)
	Reset()
}

type Result struct {
	Name string
	IR   []types.IR
}
