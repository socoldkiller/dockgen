package optimizer

import (
	"dockgen/pkg/analysis/analyzer"
	"dockgen/pkg/analysis/types"
	"fmt"
)

type WorkDirOptimizer struct {
	az *analyzer.WorkDirAnalyzer
	g  types.CFGraph
}

func NewWorkDirOptimizer(az *analyzer.WorkDirAnalyzer, g types.CFGraph) *WorkDirOptimizer {
	return &WorkDirOptimizer{az: az, g: g}
}

func (w *WorkDirOptimizer) Optimize() (Result, error) {
	f := &analyzer.FamilyAnalyzer{}
	w.az.Init(f)

	res, err := w.az.Analyze(w.g)

	if err != nil {
		return nil, err
	}

	path := mergeAdjacentPath(res)

	for _, p := range path {
		fmt.Println(p.OldIR.Cmd())
	}

	return nil, nil
}

// adjacent path
func mergeAdjacentPath(result analyzer.ResultReader) []analyzer.Item {
	var stack []analyzer.Item
	var path []analyzer.Item
	for _, res := range result.Items() {
		if len(stack) != 0 {
			top := stack[len(stack)-1]
			if top.OldIR.ID()+1 != res.OldIR.ID() {
				path = append(path, top)
				stack = nil
			}
		}
		stack = append(stack, res)
	}

	if len(stack) != 0 {
		top := stack[len(stack)-1]
		path = append(path, top)
	}

	return path

}
