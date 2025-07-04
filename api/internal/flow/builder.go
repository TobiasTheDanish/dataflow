package flow

type StepConfig struct {
	InputType        OutputType
	InputDataFormat  DataFormat
	OutputType       OutputType
	OutputDataFormat DataFormat
	HttpConfig       HttpStepConfig
}

type HttpStepConfig struct {
	method  string
	apiUrl  string
	headers map[string]string
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
		panic("DATA OUTPUT TYPE NOT IMPLEMENTED YET")
	}
}

func appendJsonInputToHttpStep(fb *FlowBuilder, config StepConfig) {
	switch config.OutputDataFormat {
	case JSON:
		s := &step{
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

func appendStep(fb *FlowBuilder, step Step) {
	if fb.start == nil {
		fb.start = step
		fb.current = fb.start
	} else {
		fb.current.SetNext(step)
		fb.current = step
	}
}
