package main

import (
	"io"
)

type Recorder struct {
	r  io.Reader
	ws []io.Writer
}

// NewRecorder 创建一个命令录制器
func NewRecorder(ws []io.Writer, r io.Reader) *Recorder {
	return &Recorder{
		r:  r,
		ws: ws,
	}
}
func (r *Recorder) Start() {
	multiWriter := io.MultiWriter(r.ws...)
	go io.Copy(multiWriter, r.r)
}
