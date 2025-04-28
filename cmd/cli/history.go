/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/log"
	"dockgen/pkg/recorder"
	"dockgen/pkg/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
		historyPath := "bash_history"
		os.Create(historyPath)
		defer os.Remove(historyPath)
		co, err := genContainer.NewContainer(genContainer.Client(), genContainer.CreateBuildAndContainerOptions{
			Tag: "debugger:0.2",
			Cmd: "/bin/sh",
			Tty: true,
			Mounts: map[string]string{
				historyPath: "/root/.bash_history",
			},
			Env: []string{
				"HISTFILE=/root/.bash_history",
				"HISTSIZE=10000",
				"HISTFILESIZE=20000",
			},
			DockerFile: "./Dockerfile",
			Raw:        true,
		})

		if err != nil {
			log.Fatalf("create container error:%v", err)
		}
		c := recorder.NewContainer(co, historyPath)
		//c.BeforeRecord([]string{"export HISTFILE=/root/.bash_history", "export HISTSIZE=10000", "export HISTFILESIZE=20000"})
		co.StartIO(os.Stdin, os.Stdout)
		co.Close()
		cmdList, err := c.Record()
		if err != nil {
			log.Fatalf("record cntainer cmd error: %v", err)
		}
		fmt.Println(util.JSONf(cmdList))
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

}
