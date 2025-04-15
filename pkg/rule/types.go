package rule

type MatchType string
type RuleAction string
type FormatType string

const (
	ActionAccept  RuleAction = "accept"
	ActionReject  RuleAction = "reject"
	ActionDrop    RuleAction = "drop"
	ActionUnknown RuleAction = "unknown"
)

const (
	JSON = "json"
	YAML = "yaml"
)

type Rule interface {
	RuleAction() RuleAction
	RuleReason() string
	RuleTags() []string
	RuleCmd() string
}

type RuleActionError struct {
	Cmd    string
	Action string
	Reason string
}
