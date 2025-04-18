package util

import "github.com/hokaccha/go-prettyjson"

func JSONf(args ...any) string {
	var s string
	for _, arg := range args {
		body, err := prettyjson.Marshal(arg)
		if err != nil {
			return s
		}
		s += string(body)
	}
	return s
}
