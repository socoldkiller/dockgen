package matcher

import (
	rs "dockgen/pkg/rules"
	"fmt"
)

type RulesMatcher interface {
	Match(cmd string) (rs.Rule, error)
}

type RulesCmdMatcher struct {
	ruleSet map[string]rs.Rule
}

func (matcher *RulesCmdMatcher) Match(cmd string) (rs.Rule, error) {
	rule, ok := matcher.ruleSet[cmd]
	if !ok {
		return rule, fmt.Errorf("not found")
	}
	return rule, nil
}

func NewRulesMatcher(rules []rs.Rule) *RulesCmdMatcher {
	ruleSet := make(map[string]rs.Rule)
	for _, rule := range rules {
		cmd := rule.RuleCmd()
		ruleSet[cmd] = rule
	}

	return &RulesCmdMatcher{ruleSet: ruleSet}
}

type RuleCmdTokenMatcher struct {
	ruleSet map[string]*CmdRuleToken
}

func NewRuleCmdTokenMatcher(rules []rs.Rule) *RuleCmdTokenMatcher {
	ruleSet := make(map[string]*CmdRuleToken)
	for _, rule := range rules {
		cmdToken := NewCmdRuleToken(rule)
		ruleSet[cmdToken.cmd] = cmdToken
	}

	return &RuleCmdTokenMatcher{
		ruleSet: ruleSet,
	}
}

func (matcher *RuleCmdTokenMatcher) Match(cmd string) (rs.Rule, error) {
	matchCmdToken := NewCmdToken(cmd)
	if r, ok := matcher.ruleSet[matchCmdToken.cmd]; ok {
		//todo
		return r.rule, nil

	}
	return nil, fmt.Errorf("not found rule")

}
