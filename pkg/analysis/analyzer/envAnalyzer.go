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
	f           *FamilyAnalyzer
}

func NewEnvAnalyzer() *EnvAnalyzer {
	return &EnvAnalyzer{
		userTable:   make(map[string]string),
		systemTable: make(map[string]string),
		f:           &FamilyAnalyzer{},
	}
}

func analyzeSystemEnv(f *FamilyAnalyzer) (map[string]string, error) {
	systemNodes := f.GetFamilyCmd("env")

	envNode, ok := lo.Last(systemNodes)
	if !ok {
		return nil, fmt.Errorf("not found system command")
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

func (e *EnvAnalyzer) Analyze(graph types.CFGraph) error {
	if err := e.f.Analyze(graph); err != nil {
		return err
	}

	systemEnv, err := analyzeSystemEnv(e.f)
	if err != nil {
		return err
	}
	e.systemTable = systemEnv

	exportNodes := e.f.GetFamilyCmd("export")
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
	e.systemTable = make(map[string]string)
	e.f = &FamilyAnalyzer{}
}

func (e *EnvAnalyzer) GetEnv() map[string]string {
	return e.userTable
}
