package flow

type StepConfig struct {
	Name                 string
	InputType            OutputType
	InputDataFormat      DataFormat
	OutputType           OutputType
	OutputDataFormat     DataFormat
	HttpConfig           HttpStepConfig
	DataProcessingConfig DataProcessingStepConfig
}

type HttpStepConfig struct {
	method  string
	apiUrl  string
	headers map[string]string
}

type DataProcessingStepConfig struct {
	keyMappings map[string]string
}

type FlowBuilder struct {
	start   Step
	current Step
}

func (fb *FlowBuilder) Build() Flow {
	return &flow{
		step: fb.start,
	}
}

func (fb *FlowBuilder) AppendStep(config StepConfig) *FlowBuilder {
	switch config.InputType {
	case HTTP:
		{
			appendHttpInputStep(fb, config)
			break
		}

	case DATA:
		{
			switch config.InputDataFormat {
			case JSON:
				appendJsonInputStep(fb, config)
				break
			case CSV:
				panic("CSV DATA FORMAT NOT IMPLEMENTED YET")
			}
			break
		}
	}

	return fb
}

func appendHttpInputStep(fb *FlowBuilder, config StepConfig) {
	switch config.OutputType {
	case HTTP:
		panic("HTTP OUTPUT TYPE NOT IMPLEMENTED YET")
	case DATA:
		appendHttpInputToDataStep(fb, config)
		break
	}
}

func appendHttpInputToDataStep(fb *FlowBuilder, config StepConfig) {
	switch config.OutputDataFormat {
	case JSON:
		s := &step{
			name: config.Name,
			processor: &HttpToJsonProcessor{
				apiUrl:  config.HttpConfig.apiUrl,
				method:  config.HttpConfig.method,
				headers: config.HttpConfig.headers,
			},
		}
		appendStep(fb, s)
		break
	}
}

func appendJsonInputStep(fb *FlowBuilder, config StepConfig) {
	switch config.OutputType {
	case HTTP:
		appendJsonInputToHttpStep(fb, config)
		break
	case DATA:
		switch config.OutputDataFormat {
		case JSON:
			appendJsonToJsonStep(fb, config)
			break
		case CSV:
			panic("CSV DATA FORMAT NOT IMPLEMENTED YET")
		}

	}
}

func appendJsonInputToHttpStep(fb *FlowBuilder, config StepConfig) {
	switch config.OutputDataFormat {
	case JSON:
		s := &step{
			name: config.Name,
			processor: &JsonToHttpProcessor{
				apiUrl:  config.HttpConfig.apiUrl,
				method:  config.HttpConfig.method,
				headers: config.HttpConfig.headers,
			},
		}
		appendStep(fb, s)
		break
	}
}

func appendJsonToJsonStep(fb *FlowBuilder, config StepConfig) {
	s := &step{
		name: config.Name,
		processor: &JsonToJsonProcessor{
			keyMappings: config.DataProcessingConfig.keyMappings,
		},
	}
	appendStep(fb, s)
}

func appendStep(fb *FlowBuilder, step Step) {
	if fb.start == nil {
		fb.start = step
		fb.current = fb.start
	} else {
		fb.current.SetNext(step)
		fb.current = step
	}
}
