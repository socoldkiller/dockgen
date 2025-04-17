package command

import (
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rule"
	"fmt"
)

type ExecutionPolicy func(e Executor, cmd string) (Result, error)

var policyTable = map[rule.RuleAction]ExecutionPolicy{
	rule.ActionAccept: Accept,
	rule.ActionDrop:   Drop,
	rule.ActionReject: Reject,
}

type RuleExecutor struct {
	executor    Executor
	policies    map[rule.RuleAction]ExecutionPolicy
	ruleMatcher matcher.RuleMatcher
}

func NewRuleExecutor(m matcher.RuleMatcher, e Executor, opts ...RuleExecutorOption) *RuleExecutor {
	ruleExecutor := &RuleExecutor{
		executor:    e,
		policies:    policyTable,
		ruleMatcher: m,
	}

	for _, opt := range opts {
		opt(ruleExecutor)
	}

	return ruleExecutor
}

type RuleExecutorOption func(*RuleExecutor)

func WithPolicyOpt(p map[rule.RuleAction]ExecutionPolicy) RuleExecutorOption {
	return func(r *RuleExecutor) {
		r.policies = p
	}
}

func (r RuleExecutor) ExecuteCommand(cmd string) (Result, error) {
	matchedRule, status := r.ruleMatcher.Match(cmd)
	policy := r.policies
	var executor ExecutionPolicy
	switch status {
	case matcher.MatchOK:
		executor = policy[matchedRule.RuleAction()]
	case matcher.MatchNotMatched:
		executor = policy[rule.ActionReject]
	case matcher.MatchNotFoundRule:
		// if r not found,run default accept
		executor = policy[rule.ActionAccept]
	default:
		executor = nil
	}

	if executor == nil {
		return Result{}, fmt.Errorf("unknown match r state")
	}

	return executor(r.executor, cmd)
}

var Accept ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return e.ExecuteCommand(cmd)
}

var Drop ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return Result{
		Cmd: cmd,
	}, fmt.Errorf("command dropped")
}

var Reject ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return Result{
		Cmd: cmd,
	}, fmt.Errorf("command rejected")
}
