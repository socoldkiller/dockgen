package dockCopy

import (
	"github.com/sirupsen/logrus"
	"io"
)

func Copy(dst io.Writer, src io.Reader) {
	if n, err := io.Copy(dst, src); err != nil {
		logrus.Debugf("Copy error: %v, copied %d bytes", err, n)
	}
}
