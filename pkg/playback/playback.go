package playback

import (
	"dockgen/pkg/builtin"
	"dockgen/pkg/command"
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rule"
	"dockgen/pkg/util"
)

type PlayBack interface {
	PlayBack(cmdList []string) ([]command.Result, error)
}

type Runner struct {
	Rules    []rule.Rule
	Matcher  matcher.RuleMatcher
	Executor command.Executor
}

type Options struct {
	Rules    []rule.Rule
	Executor command.Executor
}

func (p Runner) PlayBack(cmdList []string) ([]command.Result, error) {
	var results []command.Result

	for _, r := range cmdList {
		cmdResult, _ := p.Executor.ExecuteCommand(r)
		results = append(results, cmdResult)
	}
	return results, nil
}

func NewPlayBackRunner(pb *Options) (*Runner, error) {

	// init builtin builtinRules config
	builtinRules := builtin.LoadBuiltinConfig()
	rules, _ := util.CastSlice[builtin.Rule, rule.Rule](builtinRules)

	defaultMatcher := matcher.NewRuleCmdTokenMatcher(rules)

	ruleExecutor := command.NewRuleExecutor(defaultMatcher, pb.Executor)
	return &Runner{
		rules,
		defaultMatcher,
		ruleExecutor,
	}, nil

}
