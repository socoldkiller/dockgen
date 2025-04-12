package rules

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Rules struct {
	r io.Reader
}

const (
	JSON = "json"
	YAML = "yaml"
)

func LoadFromFile(Path string) (Rules, error) {
	r, err := os.Open(Path)
	if err != nil {
		return Rules{}, err

	}
	return Rules{
		r: r,
	}, nil
}

func (rs *Rules) Parse(Format string) ([]BuiltinRule, error) {
	var (
		rules []BuiltinRule
		err   error
	)
	switch Format {

	case JSON:
		err = json.NewDecoder(rs.r).Decode(&rules)
	case YAML:
		err = yaml.NewDecoder(rs.r).Decode(&rules)
	default:
		panic("unknown parser")
	}

	if err != nil {
		return nil, err
	}

	return rules, nil

}
