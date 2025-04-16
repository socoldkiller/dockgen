package playback

import (
	genContainer "dockgen/gen-container"
	"dockgen/pkg/builtin"
	"dockgen/pkg/command"
	"dockgen/pkg/policy/matcher"
	"dockgen/pkg/rule"
	"dockgen/pkg/util"
)

type Container struct {
	*Runner
}

func NewContainer(co *genContainer.Container) (*Container, error) {
	cfg := builtin.LoadBuiltinConfig()
	rules, _ := util.CastSlice[builtin.Rule, rule.Rule](cfg)
	m := matcher.NewRuleCmdTokenMatcher(rules)
	executor, err := command.NewStreamedContainerExecutor(co.Client, co.CreateContainer.ID, co.AttachContainer.Conn)
	ruleExecutor := command.NewRuleExecutor(m, executor)
	if err != nil {
		return nil, err
	}

	r, err := NewPlayBackRunner(&Options{
		Executor: ruleExecutor,
	})

	if err != nil {
		return nil, err
	}
	return &Container{
		Runner: r,
	}, nil

}

func (c Container) PlayBack(cmdList []string) ([]command.Result, error) {
	return c.Runner.PlayBack(cmdList)
}
