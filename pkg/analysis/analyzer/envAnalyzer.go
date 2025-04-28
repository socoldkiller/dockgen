package analyzer

import (
	"dockgen/pkg/analysis"
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

type EnvAnalyzer struct {
	userTable   map[string]string
	systemTable map[string]string
}

func NewEnvAnalyzer() *EnvAnalyzer {
	return &EnvAnalyzer{
		userTable: make(map[string]string),
	}
}

func analyzeSystemEnv(graph types.IRGraph) (map[string]string, error) {
	g, ok := graph.(*types.BashGraph)
	if !ok {
		return nil, fmt.Errorf("invalid graph type")
	}
	systemNodes := g.GetCmdGroup("env", "family")

	var envNode types.IR
	if len(systemNodes) == 0 {
		return nil, fmt.Errorf("not found system command")
	}

	envNode = systemNodes[len(systemNodes)-1]
	stdout := envNode.Stdout()
	lines := strings.Split(stdout, "\n")

	res := lo.Map(lines, func(line string, index int) *command.Result {
		return &command.Result{
			Cmd: line,
		}
	})
	envBuilder := analysis.NewBashIRBuilder()
	envList, err := envBuilder.Build(res)

	if err != nil {
		return nil, err
	}

	systemEnv := make(map[string]string)
	for _, env := range envList {
		systemEnv[env.Env().EnvVariable] = env.Env().EnvValue.Value
	}

	return systemEnv, nil
}

func (e *EnvAnalyzer) Analyze(graph types.IRGraph) error {
	systemEnv, err := analyzeSystemEnv(graph)
	if err != nil {
		return err
	}
	e.systemTable = systemEnv
	g := graph.(*types.BashGraph)

	exportNodes := g.GetCmdGroup("export", "family")
	for _, node := range exportNodes {
		env := node.Env()
		var value string
		switch {
		case env == nil:
			value = ""

		case env.EnvValue.Name != "":
			if env.EnvValue.Name[0] == '$' {
				env.EnvValue.Name = env.EnvValue.Name[1:]
			}

			value = e.systemTable[env.EnvValue.Name]

		default:
			value = env.EnvValue.Value

		}
		e.userTable[env.EnvVariable] = value
	}
	return nil
}

func (e *EnvAnalyzer) Reset() {
	e.userTable = make(map[string]string)
}

func (e *EnvAnalyzer) GetEnv() map[string]string {
	return e.userTable
}
