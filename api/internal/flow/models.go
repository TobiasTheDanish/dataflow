package flow

type DataFormat int

const (
	JSON DataFormat = iota
	CSV
)

type OutputType int

const (
	HTTP OutputType = iota
	DATA
)

type Data interface {
	Data() map[string]any
	/** Specifies to format of which the Data is, e.g. json, csv, xml, etc. */
	DataFormat() DataFormat
	/** Specifies the keys used to get the data */
	Keys() []string
}

type Input interface {
	Data() Data
}

type Processor interface {
	Process(Input) Output
}

type Output interface {
	Data() Data
	HasError() bool
	Error() error
}

type ErrorOutput struct {
	err error
}

func (o *ErrorOutput) Data() Data     { return nil }
func (o *ErrorOutput) HasError() bool { return true }
func (o *ErrorOutput) Error() error   { return o.err }

type EmptyInput struct{}

func (o *EmptyInput) Data() Data     { return &EmptyData{} }
func (o *EmptyInput) HasError() bool { return false }
func (o *EmptyInput) Error() error   { return nil }

type EmptyData struct{}

func (o *EmptyData) Data() map[string]any   { return nil }
func (o *EmptyData) DataFormat() DataFormat { return -1 }
func (o *EmptyData) Keys() []string         { return nil }
