package test

import (
	"bytes"
	analysis2 "dockgen/pkg/analysis"
	analysis "dockgen/pkg/analysis/graph"
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"dockgen/pkg/log"
	"dockgen/pkg/util"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type MockOSBackendPlayBack struct {
}

func (m MockOSBackendPlayBack) PlayBack(cmdList []string) ([]*command.Result, error) {

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

type MockFileJSONPlayBack struct {
	file string
}

func (m MockFileJSONPlayBack) PlayBack(cmdList []string) ([]*command.Result, error) {
	f, err := os.Open(m.file)
	if err != nil {
		return nil, err
	}
	var res []*command.Result
	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, err

}

func TestNewAnalyzer(t *testing.T) {

	playback := &MockFileJSONPlayBack{
		file: "playback.json",
	}

	pbReader := analysis2.NewPlayBackCmdReader(playback, nil)

	analyzer := analysis2.NewAnalyzer(pbReader, analysis2.NewBashIRBuilder())

	tokens, _ := analyzer.Analyze()

	g := types.NewBashGraph(tokens)
	g.InitGraph([]types.GraphBuilder{&analysis.BaseGraphBuilder{}, &analysis.FamilyGraphBuilder{}})
	data := g.GetTagGroup(4, "family", nil)

	for _, d := range data {
		log.Infof("%s", util.JSONf(d.Cmd()))
	}

	//
	//for _, token := range tokens {
	//	log.Infof("%s", util.JSONf(token.Program()))
	//
	//}
}
