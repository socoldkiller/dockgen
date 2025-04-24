/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// irCmd represents the ir command
var irCmd = &cobra.Command{
	Use:   "ir",
	Short: "Generate intermediate IR from recorded shell commands",
	Long: `The 'ir' command parses recorded shell commands into an Intermediate Representation (IR).
This IR serves as the foundation for later analysis, graph construction, and Dockerfile generation.
It supports input from JSON files or directly from standard input.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ir called")
	},
}

func init() {
	rootCmd.AddCommand(irCmd)

}
