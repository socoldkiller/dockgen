package command

import (
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rules"
	"fmt"
)

type ExecutionPolicy func(e Executor, cmd string) (Result, error)

var policyTable = map[rules.RuleAction]ExecutionPolicy{
	rules.ActionAccept: Accept,
	rules.ActionDrop:   Drop,
	rules.ActionReject: Reject,
}

type RuleExecutor struct {
	e           Executor
	policies    map[rules.RuleAction]ExecutionPolicy
	ruleMatcher *matcher.RulesCmdMatcher
}

func NewRuleExecutor(m *matcher.RulesCmdMatcher, e Executor, opts ...RuleExecutorOption) *RuleExecutor {
	ruleExecutor := &RuleExecutor{
		e:           e,
		policies:    policyTable,
		ruleMatcher: m,
	}

	for _, opt := range opts {
		opt(ruleExecutor)
	}

	return ruleExecutor
}

type RuleExecutorOption func(*RuleExecutor)

func WithPolicyOpt(p map[rules.RuleAction]ExecutionPolicy) RuleExecutorOption {
	return func(r *RuleExecutor) {
		r.policies = p
	}
}

func (rule RuleExecutor) ExecuteCommand(cmd string) (Result, error) {
	matchedRule, err := rule.ruleMatcher.Match(cmd)
	if err != nil {
		return Result{}, err
	}
	policy := rule.policies
	executor := policy[matchedRule.RuleAction()]
	return executor(rule.e, cmd)
}

var Accept ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return e.ExecuteCommand(cmd)
}

var Drop ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return Result{}, fmt.Errorf("command dropped")
}

var Reject ExecutionPolicy = func(e Executor, cmd string) (Result, error) {
	return Result{}, fmt.Errorf("command rejected")
}
