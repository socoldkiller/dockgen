/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/dockSession"
	"dockgen/pkg/log"
	"dockgen/pkg/util"
	"fmt"
	"github.com/spf13/cobra"
)

var longUsage = `The 'playback' command replays previously recorded shell commands 
inside a Docker container, allowing you to observe and analyze command behavior 
in a clean environment. This is useful for debugging, Dockerfile generation, 
and reproducing issues from historical terminal sessions.

Example usage:
  dockgen playback

This command mounts the recorded bash history into a new container,
executes each command, and outputs structured results (stdout, stderr, exit code, etc).`

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback",
	Short: "Replay recorded shell commands inside a Docker container",
	Long:  longUsage,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(longUsage)
		session := dockSession.NewSession("bash_history", "debugger:0.2", "./Dockerfile", genContainer.Client())
		result, err := session.Run()
		if err != nil {
			log.Fatalf("create dock session error: %v", err)
		}
		fmt.Println(util.JSONf(result))
	},
}

func init() {
	rootCmd.AddCommand(playbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playbackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playbackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
