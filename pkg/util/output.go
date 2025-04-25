package util

import (
	"dockgen/pkg/log"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func output(result interface{}, override string, outputFormat string) string {
	var jsonData string
	switch outputFormat {
	case "json":
		jsonData = JSONf(result)
	case "json-line":
		jsonData = JSONf(result)
	case "yaml":
		jsonBytes, err := yaml.Marshal(result)
		if err != nil {
			log.Fatalf("failed to unmarshal output")
		}
		jsonData = string(jsonBytes)
	default:
		return override
	}

	return jsonData
}

// SuccessOutput prints the result to stdout and exits with status code 0.
func SuccessOutput(result interface{}, override string, outputFormat string) {
	fmt.Println(output(result, override, outputFormat))
}

// ErrorOutput prints an error message to stderr and exits with status code 1.
func ErrorOutput(errResult error, override string, outputFormat string) {
	type errOutput struct {
		Error string `json:"error"`
	}
	fmt.Fprintf(os.Stderr, "%s\n", output(errOutput{errResult.Error()}, override, outputFormat))
}
