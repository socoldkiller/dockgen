package recorder

import (
	"os"
	"strings"
)

type Recorder interface {
	Record() ([]string, error)
}

type HistoryFileRecorder struct {
	path string
}

func NewHistoryFileRecorder(path string) *HistoryFileRecorder {
	return &HistoryFileRecorder{path: path}
}

func (h HistoryFileRecorder) Record() ([]string, error) {
	historyBodyByte, err := os.ReadFile(h.path)
	if err != nil {
		return nil, err
	}
	body := string(historyBodyByte)
	return strings.Split(body, "\n"), nil
}
