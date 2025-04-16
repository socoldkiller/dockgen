package builtin

import (
	"dockgen/pkg/rule"
	"embed"
	_ "embed"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
)

type Rule struct {
	Cmd    string          `json:"cmd"`
	Action rule.RuleAction `json:"action"`
}

func (r Rule) RuleCmd() string {
	return r.Cmd
}

func (r Rule) RuleAction() rule.RuleAction {
	return r.Action
}

func (r Rule) RuleReason() string {
	return ""
}

func (r Rule) RuleTags() []string {
	return nil
}

func loadRuleConfig(reader io.Reader, formatType rule.FormatType) ([]Rule, error) {
	var (
		builtinRules []Rule
		err          error
	)

	switch formatType {

	case rule.YAML:
		err = json.NewDecoder(reader).Decode(&builtinRules)

	case rule.JSON:
		err = yaml.NewDecoder(reader).Decode(&builtinRules)
	}
	return builtinRules, err
}

//go:embed rule.json
var path embed.FS

func LoadBuiltinConfig() []Rule {
	cfg, _ := path.Open("rule.json")
	rules, _ := loadRuleConfig(cfg, rule.JSON)
	return rules
}
