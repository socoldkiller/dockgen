package analyzer

import "dockgen/pkg/analysis/types"

type FamilyAnalyzer struct {
	familyNodes map[string][]types.IR
}

func (f *FamilyAnalyzer) Init(...Analyzer) {
	return
}

func (f *FamilyAnalyzer) Analyze(graph types.CFGraph) (*Result, error) {
	familyNodes := make(map[string][]types.IR)
	for _, cmd := range graph.CmdList() {
		familyNodes[cmd.Program()] = append(familyNodes[cmd.Program()], cmd)
	}
	f.familyNodes = familyNodes
	return &Result{
		Name: "FamilyAnalyzer",
	}, nil
}

func (f *FamilyAnalyzer) Reset() {
	f.familyNodes = make(map[string][]types.IR)
}

func (f *FamilyAnalyzer) GetFamilyCmd(s string) []types.IR {
	return f.familyNodes[s]
}
