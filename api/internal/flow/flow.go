package flow

type Step interface {
	Next() Step
	Start() error
	SetInput(Input)
}

type flow struct {
	Step Step
}

type Flow interface {
	Run() error
}
