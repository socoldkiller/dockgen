package matcher

import (
	"dockgen/pkg/rule"
	"fmt"
	"regexp"
	"strings"
)

type CmdToken struct {
	cmd     string
	cmdArgs string
}

func newCmdToken(cmd string) *CmdToken {
	list := strings.SplitN(cmd, " ", 2)

	if len(list) < 2 {
		list = append(list, "")
	}
	return &CmdToken{
		cmd:     strings.TrimSpace(list[0]),
		cmdArgs: strings.TrimSpace(list[1]),
	}
}

type cmdRuleToken struct {
	*CmdToken
	rule       rule.Rule
	pattern    *regexp.Regexp
	patternStr string
}

func newCmdRuleToken(r rule.Rule) *cmdRuleToken {
	cmd := r.RuleCmd()
	cmdToken := newCmdToken(cmd)
	cmd = fmt.Sprintf("%s %s", cmdToken.cmd, cmdToken.cmdArgs)
	patternStr := regexp.QuoteMeta(cmd)
	patternStr = "^" + strings.ReplaceAll(patternStr, `\*`, `.*`) + "$"
	re := regexp.MustCompile(patternStr)
	return &cmdRuleToken{
		CmdToken:   cmdToken,
		rule:       r,
		pattern:    re,
		patternStr: patternStr,
	}
}
