package emitter

import (
	"dockgen/pkg/analysis/analyzer"
	"dockgen/pkg/analysis/types"
	"fmt"
	"github.com/samber/lo"
)

type Emitter interface {
	Emit(a analyzer.Analyzer) []string
}

type EnvEmitter struct {
	g types.CFGraph
}

func (e *EnvEmitter) Emit(a analyzer.Analyzer) []string {
	res, err := a.Analyze(e.g)
	if err != nil {
		return nil
	}

	return lo.Map(res.NewIR, func(ir types.IR, index int) string {
		k := ir.Env().EnvVariable
		v := ir.Env().EnvValue.Value
		return fmt.Sprintf("%s %s=%s", types.ENV, k, v)
	})

}

func NewEnvEmitter(g types.CFGraph) *EnvEmitter {
	return &EnvEmitter{
		g: g,
	}
}
