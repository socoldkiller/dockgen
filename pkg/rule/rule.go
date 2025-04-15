package rule

type BuiltinRule struct {
	Cmd       string     `json:"cmd"`
	MatchType MatchType  `json:"matchType"`
	Pattern   string     `json:"pattern"`
	Action    RuleAction `json:"action"`
	Reason    string     `json:"reason"`
	Tags      []string   `json:"tags"`
}

func (r BuiltinRule) RuleCmd() string {
	return r.Cmd
}

func (r BuiltinRule) RuleAction() RuleAction {
	return r.Action
}

func (r BuiltinRule) RuleReason() string {
	return r.Reason
}

func (r BuiltinRule) RuleTags() []string {
	return r.Tags
}
