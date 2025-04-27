package util

import (
	"github.com/samber/lo"
)

func CastSlice[Input any, Output any](input []Input) ([]Output, error) {
	return lo.Map(input, func(item Input, index int) Output {
		return any(item).(Output)
	}), nil

}

func StrOrNil(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}
