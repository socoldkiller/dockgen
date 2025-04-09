package gen_container

import (
	"io"
)

func CopyIO(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		//todo
	}
}
