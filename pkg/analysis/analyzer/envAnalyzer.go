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
	f *FamilyAnalyzer
}

func NewEnvAnalyzer(f *FamilyAnalyzer) *EnvAnalyzer {
	if f == nil {
		f = &FamilyAnalyzer{}
	}
	return &EnvAnalyzer{
		f: f,
	}
}

func analyzeSystemEnv(f *FamilyAnalyzer) (map[string]string, error) {
	systemNodes := f.GetFamilyCmd("env")

	envNode, ok := lo.Last(systemNodes)
	if !ok {
		return nil, fmt.Errorf("not found env command, we must need env command to analyze system env")
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
		return nil, err
	}

	systemEnv := make(map[string]string)
	for _, env := range envList {
		systemEnv[env.Env().EnvVariable] = env.Env().EnvValue.Value
	}

	return systemEnv, nil
}

func (e *EnvAnalyzer) Analyze(graph types.CFGraph) (*Result, error) {
	if err := e.f.Analyze(graph); err != nil {
		return nil, err
	}
	systemEnv, err := analyzeSystemEnv(e.f)
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
		Name: "EnvAnalyzer",
		IR:   bashIRs,
	}, nil
}

func (e *EnvAnalyzer) Reset() {
	e.f = &FamilyAnalyzer{}
}
