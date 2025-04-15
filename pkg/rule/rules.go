package rule

import (
	"encoding/json"
	"os"
)

type Rules struct {
	rules []Rule
}

func ParseFromFile(Path string, formatType FormatType) (*Rules, error) {
	f, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	var rules []Rule
	switch formatType {
	case JSON:
		err = json.NewDecoder(f).Decode(&rules)
	case YAML:
		err = json.NewDecoder(f).Decode(&rules)

	}

	if err != nil {
		return nil, err
	}
	return &Rules{
		rules: rules,
	}, nil

}

//
//type Rules struct {
//	r io.Reader
//}
//
//func LoadFromFile(P ath string) (Rules, error) {
//	r, err := os.Open(Path)
//	if err != nil {
//		return Rules{}, err
//
//	}
//	return Rules{
//		r: r,
//	}, nil
//}
//
//func (rs *Rules) Parse(Format FormatType) ([]BuiltinRule, error) {
//	var (
//		rule []BuiltinRule
//		err   error
//	)
//	switch Format {
//
//	case JSON:
//		err = json.NewDecoder(rs.r).Decode(&rule)
//	case YAML:
//		err = yaml.NewDecoder(rs.r).Decode(&rule)
//	default:
//		panic("unknown parser")
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	return rule, nil
//
//}
