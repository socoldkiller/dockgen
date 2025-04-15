package matcher

import (
	"dockgen/pkg/rule"
	"fmt"
)

type MatchStatus int

const (
	MatchOK MatchStatus = iota
	MatchNotFoundRule
	MatchNotMatched
	MatchInvalidPattern
)

func (s MatchStatus) String() string {
	switch s {
	case MatchOK:
		return "MatchOK"
	case MatchNotFoundRule:
		return "MatchNotFoundRule"
	case MatchNotMatched:
		return "MatchNotMatched"
	case MatchInvalidPattern:
		return "MatchInvalidPattern"
	default:
		return "UnknownMatchStatus"
	}
}

type RuleMatcher interface {
	Match(cmd string) (rule.Rule, MatchStatus)
}

type RuleCmdMatcher struct {
	ruleSet map[string]*cmdRuleToken
}

func NewRuleCmdTokenMatcher(rules []rule.Rule) *RuleCmdMatcher {
	ruleSet := make(map[string]*cmdRuleToken)
	for _, r := range rules {
		cmdToken := newCmdRuleToken(r)
		ruleSet[cmdToken.cmd] = cmdToken
	}

	return &RuleCmdMatcher{
		ruleSet: ruleSet,
	}
}

func (matcher *RuleCmdMatcher) Match(cmd string) (rule.Rule, MatchStatus) {
	matchCmdToken := newCmdToken(cmd)
	if r, ok := matcher.ruleSet[matchCmdToken.cmd]; ok {
		return TokenMatch(matchCmdToken, r)
	}
	return nil, MatchNotFoundRule

}

func TokenMatch(cmdToken *CmdToken, matchToken *cmdRuleToken) (rule.Rule, MatchStatus) {
	cmd := fmt.Sprintf("%s %s", cmdToken.cmd, cmdToken.cmdArgs)
	if matchToken.pattern.MatchString(cmd) {
		return matchToken.rule, MatchOK
	}
	return nil, MatchNotMatched
}
