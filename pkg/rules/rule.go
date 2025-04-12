package rules

type BuiltinRule struct {
	Name      string
	MatchType MatchType
	Pattern   string
	Action    RuleAction
	Reason    string
	Tags      []string
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
