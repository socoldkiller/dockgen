package analyzer

import (
	"dockgen/pkg/analysis/types"
)

type Analyzer interface {
	Analyze(graph types.CFGraph) error
	Reset()
}
