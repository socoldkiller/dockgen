package command

type Executor interface {
	ExecuteCommand(cmd string) (Result, error)
}
