package flow

type DataFormat int

const (
	JSON DataFormat = iota
	CSV
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

type HttpData struct {
	Ok          bool
	StatusCode  int
	ContentType string
	Body        string
	data        map[string]any
	keys        []string
}

func (d HttpData) Data() map[string]any {
	if d.data == nil {
		d.data = map[string]any{
			"ok":          d.Ok,
			"statusCode":  d.StatusCode,
			"contentType": d.ContentType,
			"body":        d.Body,
		}
	}
	return d.data
}

func (d HttpData) DataFormat() DataFormat { return JSON }
func (d HttpData) Keys() []string {
	if d.keys == nil {
		d.keys = make([]string, 0, len(d.data))
		for k := range d.data {
			d.keys = append(d.keys, k)
		}
	}

	return d.keys
}

type HttpOutput struct {
	data HttpData
}

func (o *HttpOutput) Data() Data {
	return o.data
}
func (o *HttpOutput) HasError() bool { return false }
func (o *HttpOutput) Error() error   { return nil }
