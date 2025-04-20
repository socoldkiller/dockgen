package test

import (
	"bytes"
	analysis2 "dockgen/pkg/analysis"
	analysis "dockgen/pkg/analysis/graph"
	"dockgen/pkg/command"
	"dockgen/pkg/log"
	"dockgen/pkg/util"
	"os/exec"
	"strings"
	"testing"
	"time"
)

type MockPlayBack struct {
}

func (m MockPlayBack) PlayBack(cmdList []string) ([]*command.Result, error) {

	var res []*command.Result
	for _, c := range cmdList {
		var (
			stdoutBuf bytes.Buffer
			stderrBuf bytes.Buffer
		)
		cs := strings.Fields(c)
		e := exec.Command(cs[0], cs[1:]...)
		e.Stderr = &stderrBuf
		e.Stdout = &stdoutBuf
		err := e.Run()
		if err != nil {
			stderrBuf.WriteString(err.Error())
		}

		res = append(res, &command.Result{
			Cmd:    c,
			Stdout: stdoutBuf.String(),
			Stderr: stderrBuf.String(),
		})
	}
	return res, nil
}

func TestNewAnalyzer(t *testing.T) {

	playback := &MockPlayBack{}

	pbReader := analysis2.NewPlayBackCmdReader(playback, []string{"docker version", "pwd", "docker ps -a", "brew -h", "docker version"})

	analyzer := analysis2.NewAnalyzer(pbReader, analysis2.NewCmdIRBuilder())

	tokens, _ := analyzer.Analyze()

	graph := analysis.NewIRGraph(tokens)
	_ = graph.Build([]analysis.GraphBuilder{analysis.NewFamilyGraphBuilder()})
	_ = graph.Init()

	data := graph.GetTagGroup(1, "family")

	log.Infof("%s", util.JSONf(data))

	time.Sleep(1 * time.Second)
}
