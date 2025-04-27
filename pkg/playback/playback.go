package playback

import (
	"dockgen/pkg/builtin"
	"dockgen/pkg/command"
	"dockgen/pkg/log"
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rule"
	"dockgen/pkg/util"
)

type PlayBack interface {
	BeforePlayBack(cmdList []string)
	PlayBack(cmdList []string) ([]*command.Result, error)
}

type Runner struct {
	Rules    []rule.Rule
	Matcher  matcher.RuleMatcher
	Executor command.Executor
}

func (p Runner) BeforePlayBack(cmdList []string) {
	//TODO
}

type Options struct {
	Rules    []rule.Rule
	Executor command.Executor
}

func (p Runner) PlayBack(cmdList []string) ([]*command.Result, error) {
	var results []*command.Result

	for _, cmd := range cmdList {

		log.Infof("%s command start execute", cmd)
		cmdResult, _ := p.Executor.ExecuteCommand(cmd)
		log.Infof("cmd result %s", util.JSONf(cmdResult))

		results = append(results, &cmdResult)
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
