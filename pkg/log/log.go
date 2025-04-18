package log

import (
	"github.com/sirupsen/logrus"
)

type p = func(format string, args ...interface{})

func Infof(format string, args ...any) {
	logrus.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	logrus.Warnf(format, args...)
}

func Fatalf(format string, args ...any) {
	logrus.Fatalf(format, args...)
}

func Debugf(format string, args ...any) {
	logrus.Debugf(format, args...)
}
