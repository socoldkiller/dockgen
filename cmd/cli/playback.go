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

// playbackCmd represents the playback command
var playbackCmd = &cobra.Command{
	Use:   "playback",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
