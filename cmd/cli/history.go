/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/copy"
	"dockgen/pkg/playback"
	"dockgen/pkg/recorder"
	"dockgen/pkg/runner"
	"encoding/json"
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
			fmt.Println(err)
			return
		}

		go copy.Copy(co.AttachContainer.Conn, os.Stdin)
		copy.Copy(os.Stdout, co.AttachContainer.Reader)
		co.Close()
		c := recorder.NewContainer(co, historyPath)

		co1, err := genContainer.NewContainer(genContainer.Client(), genContainer.CreateBuildAndContainerOptions{
			Tag: "debugger:0.2",
			Cmd: "/bin/sh",
			Tty: false,
			Mounts: map[string]string{
				historyPath: "/root/.bash_history",
			},
			Env: []string{
				"HISTFILE=/root/.bash_history",
				"HISTSIZE=10000",
				"HISTFILESIZE=20000",
			},
			DockerFile: "./Dockerfile",
		})

		pb, err := playback.NewContainer(co1)

		if err != nil {
			return
		}

		r := runner.NewDockRunner(c, pb)
		data, err := r.Run()

		if err != nil {
			return
		}

		json.NewEncoder(os.Stdout).Encode(&data)

	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

}
