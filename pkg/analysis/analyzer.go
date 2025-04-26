package analysis

import (
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type Analyzer interface {
	Analyze(graph types.IRGraph) error
	Reset()
}

type EnvAnalyzer struct {
	envTable map[string]string
}

func NewEnvAnalyzer() *EnvAnalyzer {
	return &EnvAnalyzer{
		envTable: make(map[string]string),
	}
}

func (e *EnvAnalyzer) Analyze(graph types.IRGraph) error {
	g, ok := graph.(*types.BashGraph)
	if !ok {
		return fmt.Errorf("invalid graph type")
	}
	exportNodes := g.GetCmdGroup("env", "family")

	var envNode types.IR
	if len(exportNodes) == 0 {
		return nil
	}

	envNode = exportNodes[len(exportNodes)-1]
	stdout := envNode.Stdout()
	lines := strings.Split(stdout, "\n")

	res := lo.Map(lines, func(line string, index int) *command.Result {
		return &command.Result{
			Cmd: line,
		}
	})
	envBuilder := NewBashIRBuilder()
	envList, err := envBuilder.Build(res)
	if err != nil {
		return err
	}

	for _, env := range envList {
		e.envTable[env.Env().EnvVariable] = env.Env().EnvValue.Value
	}

	return nil
}

func (e *EnvAnalyzer) Reset() {
	e.envTable = make(map[string]string)
}

func (e *EnvAnalyzer) GetEnv() map[string]string {
	return e.envTable
}
