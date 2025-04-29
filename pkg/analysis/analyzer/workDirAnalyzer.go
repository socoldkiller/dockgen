package analyzer

import (
	"dockgen/pkg/analysis"
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"fmt"
	"github.com/samber/lo"
	"path/filepath"
)

type WorkDirAnalyzer struct {
	f   *FamilyAnalyzer
	cwd string
}

func (w *WorkDirAnalyzer) Reset() {
	return
}

func NewWorkDirAnalyzer(f *FamilyAnalyzer, initialCwd string) *WorkDirAnalyzer {
	if f == nil {
		f = &FamilyAnalyzer{}
	}
	return &WorkDirAnalyzer{
		f:   f,
		cwd: initialCwd,
	}
}

func resolvePath(base, target string) string {
	if filepath.IsAbs(target) {
		return filepath.Clean(target)
	}
	return filepath.Clean(filepath.Join(base, target))
}

func (w *WorkDirAnalyzer) Analyze(graph types.CFGraph) (*Result, error) {
	f := w.f
	if err := f.Analyze(graph); err != nil {
		return nil, err
	}

	cdNodes := f.GetFamilyCmd("cd")

	var cmdList []string

	for _, node := range cdNodes {
		if len(node.Args()) > 0 {
			target := node.Args()[0]
			w.cwd = resolvePath(w.cwd, target)
			cmd := fmt.Sprintf("cd %s", w.cwd)
			cmdList = append(cmdList, cmd)
		}
	}

	results := lo.Map(cmdList, func(cmd string, _ int) *command.Result {
		return &command.Result{
			Cmd: cmd,
		}
	})

	builder := analysis.NewBashIRBuilder()
	cdIRs, err := builder.Build(results)
	if err != nil {
		return nil, err
	}

	return &Result{
		Name:  "WorkDirAnalyzer",
		OldIR: cdNodes,
		NewIR: cdIRs,
	}, nil
}
