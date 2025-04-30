package analyzer

import "dockgen/pkg/analysis/types"

type Result struct {
	Name  string
	OldIR []types.IR
	NewIR []types.IR
	ir    map[types.IR]types.IR
}
