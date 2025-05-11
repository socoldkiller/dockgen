package util

import (
	"bytes"
	"io"
)

type DelimitedReader struct {
	r io.Reader
}

func NewDelimitedReader(r io.Reader) *DelimitedReader {
	return &DelimitedReader{
		r: r,
	}
}

func (dr *DelimitedReader) Read(p []byte) (int, error) {
	return dr.r.Read(p)
}

func (dr *DelimitedReader) ReadString(delim []byte) (string, error) {
	var buf []byte
	b := make([]byte, 1024)
	for {

		if idx := bytes.Index(buf, delim); idx != -1 {
			return string(buf[:idx]), nil
		}

		n, err := dr.Read(b)
		if err != nil {
			if err == io.EOF {
				err = nil
				continue
			}
			buf = append(buf, b[:n]...)
			return string(buf), err
		}
		buf = append(buf, b[:n]...)
	}
}
