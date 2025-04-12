package rules

type MatchType string
type RuleAction string

const (
	ActionAccept RuleAction = "accept"
	ActionReject RuleAction = "reject"
	ActionDrop   RuleAction = "drop"
)

type Rule interface {
	RuleAction() RuleAction
	RuleReason() string
	RuleTags() []string
}

type RuleActionError struct {
	Action string
	Reason string
}
