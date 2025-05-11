package command

import (
	"bytes"
	"dockgen/pkg/builtin"
	"dockgen/pkg/rule"
	"dockgen/pkg/util"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type PipeCommandExecutor struct {
	delim        string
	stdin        io.WriteCloser
	stdout       *util.DelimitedReader
	stderr       *util.DelimitedReader
	cmd          *exec.Cmd
	builtinRules map[string]builtin.Rule
}

func NewPipeCommandExecutor(CMD string, opts ...PipeExecutorOptions) (*PipeCommandExecutor, error) {
	var (
		err          error
		cmd          *exec.Cmd
		stdin        io.WriteCloser
		stderr       io.ReadCloser
		stdout       io.ReadCloser
		startCmdFunc func(startedChan chan<- error)
	)
	startCmdFunc = func(startedChan chan<- error) {

		startedChan <- cmd.Start()
		if err := cmd.Wait(); err != nil {
			log.Printf("cmd wait closed")
		}
	}

	cmd = exec.Command(CMD)

	if stdin, err = cmd.StdinPipe(); err != nil {
		return nil, err
	}

	if stdout, err = cmd.StdoutPipe(); err != nil {
		return nil, err
	}

	if stderr, err = cmd.StderrPipe(); err != nil {
		return nil, err
	}
	startedChan := make(chan error)
	go startCmdFunc(startedChan)
	if err = <-startedChan; err != nil {
		return nil, err
	}
	delim := "__CMD_DONE__"
	stdoutReader := util.NewDelimitedReader(stdout)
	stderrReader := util.NewDelimitedReader(stderr)

	p := &PipeCommandExecutor{
		stdin:        stdin,
		stdout:       stdoutReader,
		stderr:       stderrReader,
		delim:        delim,
		cmd:          cmd,
		builtinRules: nil,
	}

	for _, opt := range opts {
		if err = opt(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (pipe *PipeCommandExecutor) ExecuteCommand(cmd string) (Result, error) {
	r, found := pipe.builtinRules[cmd]
	if !found {
		// ok,we use "" name to instead of not found rule
		if r, found = pipe.builtinRules[""]; !found {
			return Result{}, rule.NewRejectError(cmd, fmt.Sprintf("%s rule not found", cmd))
		}

	}

	switch r.RuleAction() {
	case rule.ActionAccept:
		return pipe.executeCommand(cmd)

	case rule.ActionDrop:

		return Result{
			Cmd: cmd,
		}, rule.NewDropError(cmd, "blocked by rule: command dropped")

	case rule.ActionReject:
		return Result{
			Cmd: cmd,
		}, rule.NewRejectError(cmd, "blocked by rule: command rejected")

	}

	return Result{}, fmt.Errorf("unknown rule action")

}

func (pipe *PipeCommandExecutor) executeCommand(cmd string) (Result, error) {
	var (
		outBuf = new(bytes.Buffer)
		errBuf = new(bytes.Buffer)
		delim  = pipe.delim
		err    error
	)

	fullCmd := fmt.Sprintf("%s; echo %s; echo %s 1>&2\n", cmd, delim, delim)

	if err != nil {
		return Result{}, err
	}

	if _, err = io.Copy(pipe.stdin, strings.NewReader(fullCmd)); err != nil {
		return Result{}, err
	}

	if _, err = io.Copy(outBuf, pipe.stdout); err != nil {
		return Result{}, err
	}

	if _, err = io.Copy(errBuf, pipe.stderr); err != nil {
		return Result{}, err
	}

	res := Result{
		Cmd:    cmd,
		Stdout: outBuf.String(),
		Stderr: errBuf.String(),
	}
	return res, nil
}

type PipeExecutorOptions = func(*PipeCommandExecutor) error

func WithRuleFile(path string, FormatType rule.FormatType) PipeExecutorOptions {
	return func(executor *PipeCommandExecutor) error {
		return nil
	}
}
