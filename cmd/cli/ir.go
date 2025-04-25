/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"dockgen/pkg/analysis"
	"dockgen/pkg/analysis/types"
	"dockgen/pkg/command"
	"dockgen/pkg/util"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
)

type ExportIR struct {
	Cmd         string            `json:"cmd,omitempty"`
	Program     string            `json:"program,omitempty"`
	PipeCommand []*ExportIR       `json:"pipeCommand,omitempty"`
	Options     map[string]string `json:"options,omitempty"`
	Args        []string          `json:"args,omitempty"`
	Input       *string           `json:"input,omitempty"`
	Output      *string           `json:"output,omitempty"`
	Redirect    *string           `json:"redirect,omitempty"`
	Env         *types.Env        `json:"env,omitempty"`
}

func NewExportIR(ir types.IR) *ExportIR {
	var pipeExportIR []*ExportIR
	for _, pipe := range ir.PipeCommand() {
		ir := NewExportIR(pipe)
		pipeExportIR = append(pipeExportIR, ir)
	}

	return &ExportIR{
		Cmd:         ir.Cmd(),
		Program:     ir.Program(),
		Options:     ir.Options(),
		Args:        ir.Args(),
		Input:       ir.Input(),
		Output:      ir.Output(),
		Redirect:    ir.Redirect(),
		Env:         ir.Env(),
		PipeCommand: pipeExportIR,
	}

}

// irCmd represents the ir command
var irCmd = &cobra.Command{
	Use:   "ir",
	Short: "Generate intermediate IR from recorded shell commands",
	Long: `The 'ir' command parses recorded shell commands into an Intermediate Representation (IR).
This IR serves as the foundation for later analysis, graph construction, and Dockerfile generation.
It supports input from JSON files or directly from standard input.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var IRList []*ExportIR
		for _, arg := range args {
			f, err := os.Open(arg)

			if err != nil {
				continue
			}

			var result []*command.Result
			err = json.NewDecoder(f).Decode(&result)
			if err != nil {
				continue
			}

			builder := analysis.NewBashIRBuilder()
			cmdIRs, err := builder.Build(result)
			if err != nil {
				continue
			}

			for _, ir := range cmdIRs {
				export := NewExportIR(ir)
				IRList = append(IRList, export)
			}
		}

		util.SuccessOutput(IRList, "", "json")
	},
}

func init() {
	rootCmd.AddCommand(irCmd)

}
