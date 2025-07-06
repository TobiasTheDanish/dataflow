package flow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JsonData struct {
	data map[string]any
	keys []string
}

func (d *JsonData) Data() map[string]any {
	return d.data
}
func (d *JsonData) DataFormat() DataFormat {
	return JSON
}
func (d *JsonData) Keys() []string {
	if d.keys == nil {
		d.keys = make([]string, 0, len(d.data))
		for k := range d.data {
			d.keys = append(d.keys, k)
		}
	}

	return d.keys
}

type JsonInput struct {
	rawData    []byte
	parsedData map[string]any
}

func (i *JsonInput) Data() Data {
	if i.parsedData == nil {
		if err := json.Unmarshal(i.rawData, &i.parsedData); err != nil {
			panic(err)
		}
	}

	return &JsonData{
		data: i.parsedData,
	}
}
func (i *JsonInput) HasError() bool { return false }
func (i *JsonInput) Error() error   { return nil }

type JsonToHttpProcessor struct {
	method  string
	apiUrl  string
	headers map[string]string
}

func (p *JsonToHttpProcessor) Process(input Input) Output {
	var req *http.Request
	data := input.Data().Data()
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return &ErrorOutput{err}
		}

		req, err = http.NewRequest(p.method, p.apiUrl, bytes.NewReader(body))
		if err != nil {
			return &ErrorOutput{err}
		}
	} else {
		request, err := http.NewRequest(p.method, p.apiUrl, nil)
		if err != nil {
			return &ErrorOutput{err}
		}
		req = request
	}

	for k, v := range p.headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &ErrorOutput{err}
	}

	if res.StatusCode == 200 || res.StatusCode == 201 {
		// Success with potential body
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return &ErrorOutput{err}
		}

		return &HttpOutput{
			data: HttpResponseData{
				Ok:          true,
				StatusCode:  res.StatusCode,
				ContentType: res.Header.Get("Content-Type"),
				Body:        string(body),
			},
		}
	} else if res.StatusCode == 202 || res.StatusCode == 204 {
		return &HttpOutput{
			data: HttpResponseData{
				Ok:          true,
				StatusCode:  res.StatusCode,
				ContentType: res.Header.Get("Content-Type"),
				Body:        "",
			},
		}
	} else {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return &ErrorOutput{err}
		}

		return &HttpOutput{
			data: HttpResponseData{
				Ok:          false,
				StatusCode:  res.StatusCode,
				ContentType: res.Header.Get("Content-Type"),
				Body:        string(body),
			},
		}
	}
}

type JsonToJsonProcessor struct {
	keyMappings map[string]string
}

func (p *JsonToJsonProcessor) Process(input Input) Output {
	if in, ok := input.(*JsonInput); ok {
		data := in.Data()
		if data.DataFormat() != JSON {
			return &ErrorOutput{err: fmt.Errorf("Expected JSON data as input")}
		}

		inputDataMap := data.Data()
		outputDataMap := make(map[string]any)
		for inKey, outKey := range p.keyMappings {
			outputDataMap[outKey] = inputDataMap[inKey]
		}

		return &JsonInput{
			parsedData: outputDataMap,
		}
	} else {
		return &ErrorOutput{err: fmt.Errorf("Expected JSON input")}
	}
}
