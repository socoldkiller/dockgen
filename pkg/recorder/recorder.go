package recorder

import (
	"os"
	"strings"
)

type Recorder interface {
	BeforeRecord(cmd []string)
	Record() ([]string, error)
	AfterRecord(cmd []string)
}

type HistoryFileRecorder struct {
	path string
}

func (h HistoryFileRecorder) BeforeRecord([]string) {
	return
}

func (h HistoryFileRecorder) AfterRecord([]string) {
	return
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
