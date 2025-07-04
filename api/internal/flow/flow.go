package flow

type step struct {
	next      Step
	processor Processor
	input     Input
}

func (s *step) Next() Step { return s.next }
func (s *step) Start() error {
	out := s.processor.Process(s.input)
	if out.HasError() {
		return out.Error()
	}

	if s.Next() != nil {
		s.Next().SetInput(out)
	}

	return nil
}
func (s *step) SetInput(in Input) { s.input = in }
func (s *step) SetNext(step Step) { s.next = step }

type Step interface {
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
		if err := f.step.Start(); err != nil {
			f.step = first
			return err
		}
		f.step = f.step.Next()
	}

	f.step = first
	return nil
}
