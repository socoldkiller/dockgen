package analyzer

import (
	"dockgen/pkg/analysis/types"
	"github.com/samber/lo"
)

type RunAnalyzer struct {
	f *FamilyAnalyzer
}

func (r *RunAnalyzer) Init(az ...Analyzer) {
	if len(az) == 0 {
		r.f = new(FamilyAnalyzer)
		return
	}
	r.f = az[0].(*FamilyAnalyzer)
}

func NewRunAnalyzer() *RunAnalyzer {
	return &RunAnalyzer{}
}

func (r *RunAnalyzer) Analyze(graph types.CFGraph) (ResultReader, error) {
	irs := lo.Filter(graph.CmdList(), func(item types.IR, index int) bool {
		return item.Program() != "export" && item.Program() != "cd"
	})

	return NewResult("RunAnalyzer", irs, irs), nil
}

func (r *RunAnalyzer) Reset() {
	return
}
