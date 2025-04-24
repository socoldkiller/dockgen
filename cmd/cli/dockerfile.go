/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerfileCmd represents the dockerfile command
var dockerfileCmd = &cobra.Command{
	Use:   "dockerfile",
	Short: "Generate a Dockerfile from IR",
	Long: `The 'dockerfile' command transforms the IR into a reproducible and optimized Dockerfile.
It intelligently handles environment variables, command ordering, and file operations to create clean build layers.
Output can be written to a file or printed to standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dockerfile called")
	},
}

func init() {
	rootCmd.AddCommand(dockerfileCmd)

}
