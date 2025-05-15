package optimizer

type Result interface {
}

type Optimizer interface {
	Optimize() (Result, error)
}
