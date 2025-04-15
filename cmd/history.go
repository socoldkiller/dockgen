/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/command"
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rand"
	"dockgen/pkg/rule"
	convert "dockgen/pkg/util"
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
	"time"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		f, err := os.Open("bash_history")
		if err != nil {
			return
		}

		historyCmd, err := HistoryCmdByFile(f)

		if err != nil {
			return
		}

		list, err := HistoryCommandStd(historyCmd)

		if err != nil {
			return
		}

		json.NewEncoder(os.Stdout).Encode(list)

		time.Sleep(2 * time.Second)

	},
}

func HistoryCmdByFile(file *os.File) ([]string, error) {
	history, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	historyStr := string(history)
	return strings.Split(historyStr, "\n"), nil
}

func HistoryCommandStd(cmds []string) ([]command.Result, error) {

	c := genContainer.Client()
	ContainerName := rand.String(5)
	attachContainer, err := genContainer.CreateAttachContainer(c,
		genContainer.CreateBuildAndContainerOptions{
			Tag:           "debugger:0.2",
			ContainerName: ContainerName,
			Cmd:           "/bin/sh",
			Tty:           false,
			Mounts: map[string]string{
				"bash_history": "/root/.bash_history",
			},
			Env: []string{
				"HISTFILE=/root/.bash_history",
				"HISTSIZE=10000",
				"HISTFILESIZE=20000",
			},
			DockerFile: "./Dockerfile",
		})

	if err != nil {
		return nil, err
	}

	executor, err := command.NewStreamedContainerExecutor(c, ContainerName, attachContainer.Conn)

	if err != nil {
		return nil, err
	}
	var f *os.File
	if f, err = os.Open("rule.json"); err != nil {
		return nil, err
	}
	var rs []rule.BuiltinRule
	if err = json.NewDecoder(f).Decode(&rs); err != nil {
		return nil, err
	}
	irs, _ := convert.CastSlice[rule.BuiltinRule, rule.Rule](rs)

	m := matcher.NewRuleCmdTokenMatcher(irs)

	ruleExecutor := command.NewRuleExecutor(m, executor)

	var results []command.Result
	for _, cmd := range cmds {
		res, err := ruleExecutor.ExecuteCommand(cmd)
		if err != nil {
			continue
		}
		results = append(results, res)
	}

	return results, nil

}

func init() {
	rootCmd.AddCommand(historyCmd)

}
