/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Build a command dependency graph from IR",
	Long: `The 'graph' command constructs a visual dependency graph based on the intermediate IR.
It reveals the logical flow and relationships between commands, which is useful for optimization and debugging.
The output can be exported in DOT format for Graphviz or rendered as an image.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("graph called")
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)

}
