package analyzer

import (
	"dockgen/pkg/analysis/types"
)

type Analyzer interface {
	Init(az ...Analyzer)
	Analyze(graph types.CFGraph) (*Result, error)
	Reset()
}

func GetResultIR(res *Result, idx int) (types.IR, types.IR) {
	return res.OldIR[idx], res.NewIR[idx]
}
