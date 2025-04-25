/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"dockgen/pkg/analysis"
	"dockgen/pkg/command"
	"dockgen/pkg/util"
	"fmt"
	"github.com/chzyer/readline"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strings"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.SuccessOutput(nil, "welcome to ir interactive terminal", "text")
		rl, err := readline.NewEx(&readline.Config{
			Prompt:          color.New(color.FgGreen).Sprintf("%s", ">>> "),
			InterruptPrompt: "^C",
			EOFPrompt:       "exit",
		})
		if err != nil {
			panic(err)
		}
		defer rl.Close()

		builder := analysis.NewBashIRBuilder()
		for {
			line, err := rl.Readline()
			if err != nil {
				break
			}
			line = strings.TrimSpace(line)
			switch line {
			case "":
				continue
			case "exit":
				return
			case "clear":
				fmt.Print("\033[H\033[2J")
				continue
			default:

			}
			result := &command.Result{Cmd: line}
			cmdIR, err := builder.Build([]*command.Result{result})
			if err != nil {
				continue
			}
			util.SuccessOutput(NewExportIR(cmdIR[0]), "", "json")
		}

	},
}

func init() {
	irCmd.AddCommand(watchCmd)
}
