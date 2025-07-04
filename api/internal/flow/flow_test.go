package flow

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFlowBuilder(t *testing.T) {
	steps := []StepConfig{
		{
			InputType:        HTTP,
			InputDataFormat:  JSON,
			OutputType:       DATA,
			OutputDataFormat: JSON,
			HttpConfig: HttpStepConfig{
				method:  "GET",
				apiUrl:  "https://jsonplaceholder.typicode.com/posts/1",
				headers: map[string]string{},
			},
		},
		{
			OutputType:       HTTP,
			OutputDataFormat: JSON,
			InputType:        DATA,
			InputDataFormat:  JSON,
			HttpConfig: HttpStepConfig{
				method: "POST",
				apiUrl: "https://jsonplaceholder.typicode.com/posts",
				headers: map[string]string{
					"Content-Type": "application/json; charset=UTF-8",
				},
			},
		},
	}

	var fb FlowBuilder

	for _, stepConfig := range steps {
		fb.AppendStep(stepConfig)
	}

	f := fb.Build()

	if err := f.Run(); err != nil {
		t.Fatalf("Flow failed with error: %v", err)
	}
}
