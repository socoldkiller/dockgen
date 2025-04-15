package convert

import (
	"fmt"
)

func CastSlice[Input any, Output any](input []Input) ([]Output, error) {
	var result []Output
	for _, item := range input {
		casted, ok := any(item).(Output)
		if !ok {
			return nil, fmt.Errorf("cast failed for item: %+v", item)
		}
		result = append(result, casted)
	}
	return result, nil
}
