package rule

import "fmt"

func (e *RuleActionError) Error() string {
	return fmt.Sprintf("command %s: %s", e.Action, e.Reason)
}

func NewDropError(cmd, reason string) error {
	return &RuleActionError{
		Cmd:    cmd,
		Action: "dropped",
		Reason: reason,
	}
}

func NewRejectError(cmd, reason string) error {
	return &RuleActionError{
		Cmd:    cmd,
		Action: "rejected",
		Reason: reason,
	}
}
