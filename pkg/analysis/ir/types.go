package ir

type Variable struct {
	Name  string
	Value string
}

type Env struct {
	EnvVariable string
	EnvValue    Variable
}

type BashCommandIR struct {
	ID          int
	Stdout      string
	Stderr      string
	Program     string
	PipeCommand []*BashCommandIR
	Options     map[string]string
	Args        []string
	Input       *string
	Output      *string
	Redir       string
	Env         *Env
}
