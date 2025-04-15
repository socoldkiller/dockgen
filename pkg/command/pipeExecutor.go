package command

import (
	"bytes"
	"dockgen/pkg/rules"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

type DelimitedReader struct {
	r      io.Reader
	delim  []byte
	buffer []byte
	eofHit bool
}

func NewDelimitedReader(r io.Reader, delim string) *DelimitedReader {
	return &DelimitedReader{
		r:     r,
		delim: []byte(delim),
	}
}

func (dr *DelimitedReader) Read(p []byte) (int, error) {
	if dr.eofHit {
		dr.eofHit = false
	}
	for {
		if idx := bytes.Index(dr.buffer, dr.delim); idx != -1 {
			n := copy(p, dr.buffer[:idx])
			dr.buffer = dr.buffer[idx+len(dr.delim):]
			dr.eofHit = true
			return n, io.EOF
		}

		tmp := make([]byte, 1024)
		n, err := dr.r.Read(tmp)
		if n > 0 {
			dr.buffer = append(dr.buffer, tmp[:n]...)
		}

		if err != nil {
			if len(dr.buffer) > 0 {
				n := copy(p, dr.buffer)
				dr.buffer = nil
				return n, io.EOF
			}
			return 0, err
		}
	}
}

func getChildrenPid(pid int) ([]int, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		return nil, err
	}

	children, err := p.Children()
	if err != nil {
		return nil, err
	}

	var childrenPid []int
	for _, child := range children {
		childrenPid = append(childrenPid, int(child.Pid))
	}
	return childrenPid, nil
}

func destroyChildrenProcess(childrenPid []int) {
	for _, pid := range childrenPid {
		_ = syscall.Kill(pid, syscall.SIGKILL)
	}
}

type PipeCommandExecutor struct {
	delim        string
	stdin        io.WriteCloser
	stdout       io.Reader
	stderr       io.Reader
	cmd          *exec.Cmd
	builtinRules map[string]rules.BuiltinRule
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
	stdoutReader := NewDelimitedReader(stdout, delim)
	stderrReader := NewDelimitedReader(stderr, delim)

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
			return Result{}, rules.NewRejectError(cmd, fmt.Sprintf("%s rule not found", cmd))
		}

	}

	switch r.RuleAction() {
	case rules.ActionAccept:
		return pipe.executeCommand(cmd)

	case rules.ActionDrop:

		return Result{}, rules.NewDropError(cmd, "blocked by rule: command dropped")

	case rules.ActionReject:
		return Result{}, rules.NewRejectError(cmd, "blocked by rule: command rejected")

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

func WithRuleFile(path string, FormatType rules.FormatType) PipeExecutorOptions {
	return func(executor *PipeCommandExecutor) error {
		return nil
	}
}
