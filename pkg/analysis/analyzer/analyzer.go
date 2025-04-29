package analyzer

import (
	"dockgen/pkg/analysis/types"
)

type Analyzer interface {
	Analyze(graph types.CFGraph) (*Result, error)
	Reset()
}

type Result struct {
	Name  string
	OldIR []types.IR
	NewIR []types.IR
}

func GetResultIR(res *Result, idx int) (types.IR, types.IR) {
	return res.OldIR[idx], res.NewIR[idx]
}
