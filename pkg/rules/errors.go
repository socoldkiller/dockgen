package rules

import "fmt"

func (e *RuleActionError) Error() string {
	return fmt.Sprintf("command %s: %s", e.Action, e.Reason)
}

func NewDropError(reason string) error {
	return &RuleActionError{
		Action: "dropped",
		Reason: reason,
	}
}

func NewRejectError(reason string) error {
	return &RuleActionError{
		Action: "rejected",
		Reason: reason,
	}
}
