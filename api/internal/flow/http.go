package flow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HttpToJsonProcessor struct {
	method  string
	apiUrl  string
	headers map[string]string
}

func (p *HttpToJsonProcessor) Process(input Input) Output {
	var bodyReader io.Reader
	if input != nil {
		data := input.Data()
		if data != nil {
			switch data.DataFormat() {
			case JSON:
				{
					body, err := json.Marshal(data.Data())
					if err != nil {
						return &ErrorOutput{fmt.Errorf("Error marshalling request body: %v", err)}
					}
					bodyReader = bytes.NewReader(body)
					break
				}
			default:
				panic(fmt.Sprintf("Data format %v not implemented", data.DataFormat()))
			}
		} else {
			bodyReader = nil
		}
	}

	req, err := http.NewRequest(p.method, p.apiUrl, bodyReader)
	if err != nil {
		return &ErrorOutput{err}
	}

	for k, v := range p.headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &ErrorOutput{err}
	}

	if res.StatusCode == 200 || res.StatusCode == 201 {
		contentType := res.Header.Get("Content-Type")
		if !strings.Contains(contentType, "application/json") {
			return &ErrorOutput{err: fmt.Errorf("Http request did not return JSON. Content type was '%s'", contentType)}
		}

		// Success with potential body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return &ErrorOutput{err}
		}
		defer res.Body.Close()

		var jsonBody map[string]any
		err = json.Unmarshal(body, &jsonBody)
		if err != nil {
			return &ErrorOutput{err: fmt.Errorf("Unable to unmarshal response body: %v", err)}
		}

		return &JsonInput{
			rawData:    body,
			parsedData: jsonBody,
		}
	} else if res.StatusCode == 202 || res.StatusCode == 204 {
		return &ErrorOutput{err: fmt.Errorf("Http request did not return JSON")}
	} else {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return &ErrorOutput{err}
		}

		return &ErrorOutput{err: fmt.Errorf("Http request failed with status %s.\nBody: %s", res.Status, body)}
	}
}

type HttpInputData struct {
	data           any
	marshalledData map[string]any
	dataFormat     DataFormat
	marshaller     HttpInputDataMarshaller
}

func (d *HttpInputData) Data() map[string]any {
	if d.marshalledData == nil {
		marshalledData, err := d.marshaller.Marshal(d.data)
		if err != nil {
			panic(err)
		}
		d.marshalledData = marshalledData
	}

	return d.marshalledData
}
func (d *HttpInputData) DataFormat() DataFormat {
	return d.dataFormat
}
func (d *HttpInputData) Keys() []string {
	keys := make([]string, 0, len(d.marshalledData))
	for k := range d.marshalledData {
		keys = append(keys, k)
	}

	return keys
}

type HttpInputDataMarshaller interface {
	Marshal(any) (map[string]any, error)
}

type HttpInput struct {
	data HttpInputData
}

func (i *HttpInput) Data() Data {
	return &i.data
}

type HttpResponseData struct {
	Ok          bool
	StatusCode  int
	ContentType string
	Body        string
	data        map[string]any
	keys        []string
}

func (d HttpResponseData) Data() map[string]any {
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

func (d HttpResponseData) DataFormat() DataFormat { return JSON }
func (d HttpResponseData) Keys() []string {
	if d.keys == nil {
		d.keys = make([]string, 0, len(d.data))
		for k := range d.data {
			d.keys = append(d.keys, k)
		}
	}

	return d.keys
}

type HttpOutput struct{ data HttpResponseData }

func (o *HttpOutput) Data() Data     { return o.data }
func (o *HttpOutput) HasError() bool { return false }
func (o *HttpOutput) Error() error   { return nil }
