package flow

import "log/slog"

type step struct {
	name      string
	next      Step
	processor Processor
	input     Input
}

func (s *step) Name() string { return s.name }
func (s *step) Next() Step   { return s.next }
func (s *step) Start() error {
	slog.Debug("Step input", "name", s.name, "input", s.input)

	out := s.processor.Process(s.input)
	if out.HasError() {
		slog.Debug("Step output with error", "name", s.name, "output", out)
		return out.Error()
	}
	slog.Debug("Step output", "name", s.name, "output", out)

	if s.Next() != nil {
		s.Next().SetInput(out)
	}

	return nil
}
func (s *step) SetInput(in Input) { s.input = in }
func (s *step) SetNext(step Step) { s.next = step }

type Step interface {
	Name() string

	Next() Step
	Start() error

	SetNext(Step)
	SetInput(Input)
}

type flow struct {
	step Step
}

type Flow interface {
	Run() error
}

/*
Runs the steps of flow by following the linked list of steps.

The flow resets itself upon return. If a step returns an error, Run returns that error.
*/
func (f *flow) Run() error {
	first := f.step

	for f.step != nil {
		slog.Info("Running step", "name", f.step.Name())

		if err := f.step.Start(); err != nil {
			slog.Info("Step failed", "name", f.step.Name(), "error", err)

			f.step = first
			return err
		}
		slog.Info("Step successful", "name", f.step.Name())
		f.step = f.step.Next()
	}

	f.step = first
	return nil
}
