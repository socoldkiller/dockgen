package analysis

import (
	"bytes"
	"dockgen/pkg/command"
	"dockgen/pkg/log"
	"dockgen/pkg/util"
	"os/exec"
	"strings"
	"testing"
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

	pbReader := NewPlayBackCmdReader(playback, []string{"ls -l", "pwd", "echo hello world", "brew -h", "docker version"})

	analyzer := NewAnalyzer(pbReader, NewCmdIRBuilder())

	tokens, _ := analyzer.Analyze()

	graph := NewIRGraph(tokens, []GraphBuilder{BaseGraphBuilder{}})

	tokens = graph.GetTagGroup(1, "context", nil)

	log.Infof("%s", util.JSONf(tokens))
}
