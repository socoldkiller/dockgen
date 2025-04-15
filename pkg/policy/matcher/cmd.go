package matcher

import (
	"dockgen/pkg/rules"
	"strings"
)

type CmdToken struct {
	cmd  string
	args []string
}

func NewCmdToken(cmd string) *CmdToken {
	list := strings.Split(cmd, " ")
	return &CmdToken{
		cmd:  list[0],
		args: list[1:],
	}
}

type CmdRuleToken struct {
	*CmdToken
	rule rules.Rule
}

func NewCmdRuleToken(r rules.Rule) *CmdRuleToken {
	cmd := r.RuleCmd()
	return &CmdRuleToken{
		CmdToken: NewCmdToken(cmd),
		rule:     r,
	}
}
