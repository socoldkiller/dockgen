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
	f           *FamilyAnalyzer
	systemNodes []types.IR
}

func (e *EnvAnalyzer) Init(az ...Analyzer) {
	if len(az) == 0 {
		e.f = new(FamilyAnalyzer)
		return
	}
	e.f = az[0].(*FamilyAnalyzer)
}

func ParseEnvCmdIR(envNode types.IR) []types.IR {
	if envNode.Cmd() != "env" {
		return nil
	}

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
		return nil
	}
	return envList
}

func analyzeSystemEnv(a *EnvAnalyzer) (map[string]string, error) {
	systemEnv := make(map[string]string)
	for _, env := range a.systemNodes {
		systemEnv[env.Env().EnvVariable] = env.Env().EnvValue.Value
	}
	return systemEnv, nil
}

func NewEnvAnalyzer(systemNodes []types.IR) *EnvAnalyzer {
	envAnalyzer := &EnvAnalyzer{
		systemNodes: systemNodes,
	}
	return envAnalyzer
}

func (e *EnvAnalyzer) Analyze(graph types.CFGraph) (*Result, error) {
	_, err := e.f.Analyze(graph)
	if err != nil {
		return nil, err
	}

	systemEnv, err := analyzeSystemEnv(e)
	if err != nil {
		return nil, err
	}

	exportNodes := e.f.GetFamilyCmd("export")
	builder := analysis.NewBashIRBuilder()
	var exportCmdList []string
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

			value = systemEnv[env.EnvValue.Name]

		default:
			value = env.EnvValue.Value

		}
		exportCmd := fmt.Sprintf("export %s=%s", env.EnvVariable, value)
		exportCmdList = append(exportCmdList, exportCmd)
	}

	cmdResults := lo.Map(exportCmdList, func(cmd string, _ int) *command.Result {
		return &command.Result{
			Cmd: cmd,
		}
	})

	bashIRs, err := builder.Build(cmdResults)
	if err != nil {
		return nil, err
	}

	return &Result{
		Name:  "EnvAnalyzer",
		OldIR: exportNodes,
		NewIR: bashIRs,
	}, nil
}

func (e *EnvAnalyzer) Reset() {
	e.f = &FamilyAnalyzer{}
}
