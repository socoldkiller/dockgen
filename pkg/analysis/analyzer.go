package analysis

import (
	"dockgen/pkg/analysis/ir"
	"fmt"
)

type Analyzer struct {
	cmdReader CmdReader
	builder   IRBuilder
}

func NewAnalyzer(cmdReader CmdReader, builder IRBuilder) *Analyzer {
	return &Analyzer{
		cmdReader: cmdReader,
		builder:   builder,
	}
}

func (a *Analyzer) Analyze() ([]*ir.BashCommandIR, error) {
	commands, err := a.cmdReader.Read()
	if err != nil {
		return nil, fmt.Errorf("read commands: %w", err)
	}

	irs, err := a.builder.Build(commands)
	if err != nil {
		return nil, fmt.Errorf("build BashCommandIR: %w", err)
	}

	return irs, nil
}
