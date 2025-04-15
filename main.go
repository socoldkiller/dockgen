/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"dockgen/cmd"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"time"
)

func setupLogLevel() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	switch strings.ToLower(logLevel) {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel) // 默认设置为 info
	}
}

func setupLogRus() {
	setupLogLevel()
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.DateTime,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			funcName := f.Function
			short := funcName[strings.LastIndex(funcName, ".")+1:]
			return short, f.File + ":" + fmt.Sprint(f.Line)
		},
		ForceColors: true,
	})
	logrus.SetReportCaller(true)
}

func main() {
	setupLogRus()
	cmd.Execute()
}
