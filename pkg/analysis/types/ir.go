package types

type IR interface {
	Cmd() string
	ID() int
	Stdout() string
	Stderr() string
	Program() string
	PipeCommand() []IR
	Options() map[string]string
	Args() []string
	Input() *string
	Output() *string
	Redirect() *string
	Env() *Env
}
