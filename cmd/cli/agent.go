/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"dockgen/pkg/command"
	"dockgen/pkg/rule"
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

func AgentCommandHandler() (func(w http.ResponseWriter, r *http.Request), error) {

	executor, err := command.NewPipeCommandExecutor("bash",
		command.WithRuleFile("rule.json", rule.JSON))
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		c := string(body)
		if c == "" {
			json.NewEncoder(w).Encode(command.Result{})
			return
		}
		w.Header().Set("Content-Type", "application/json")

		res, err := executor.ExecuteCommand(c)
		switch err {
		case nil:
			json.NewEncoder(w).Encode(res)

		default:
			json.NewEncoder(w).Encode(err)

		}

	}, nil

}

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		commandHandler, err := AgentCommandHandler()
		if err != nil {
			log.Fatal(err)
		}
		http.HandleFunc("/", commandHandler)
		log.Println("port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))

	},
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
