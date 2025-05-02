package analyzer

import (
	"dockgen/pkg/analysis"
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"fmt"
	"github.com/samber/lo"
	"path/filepath"
	"strings"
)

type WorkDirAnalyzer struct {
	f   *FamilyAnalyzer
	cwd string
}

func (w *WorkDirAnalyzer) Init(az ...Analyzer) {
	if len(az) == 0 {
		w.f = new(FamilyAnalyzer)
		return
	}
	w.f = az[0].(*FamilyAnalyzer)
}

func (w *WorkDirAnalyzer) Reset() {
	return
}

func NewWorkDirAnalyzer(initialCwd string) *WorkDirAnalyzer {
	return &WorkDirAnalyzer{
		cwd: initialCwd,
	}
}

func resolvePath(base, target string) string {
	if filepath.IsAbs(target) {
		return filepath.Clean(target)
	}
	return filepath.Clean(filepath.Join(base, target))
}

func isValidCD(node types.IR) bool {
	stderr := node.Stderr()
	if stderr == "" {
		return true
	}
	if strings.Contains(stderr, "No such file or directory") ||
		strings.Contains(stderr, "cd:") {
		return false
	}
	return true
}

func (w *WorkDirAnalyzer) Analyze(graph types.CFGraph) (ResultReader, error) {
	f := w.f
	if _, err := f.Analyze(graph); err != nil {
		return nil, err
	}

	cdNodes := f.GetFamilyCmd("cd")

	cdNodes = lo.Filter(cdNodes, func(cdIR types.IR, _ int) bool {
		return isValidCD(cdIR)
	})

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
	return NewResult("WorkDirAnalyzer", cdNodes, cdIRs), nil
}
