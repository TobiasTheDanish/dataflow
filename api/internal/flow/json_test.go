package flow

import (
	"reflect"
	"testing"
)

type ProcessorTest struct {
	name string
	p    Processor
	in   Input
	out  Output
}

func TestJsonToJsonProcessors(t *testing.T) {
	tests := []ProcessorTest{
		{
			name: "Test simple happy path",
			p: &JsonToJsonProcessor{
				keyMappings: map[string]string{
					"name": "Name",
					"age":  "Age",
				},
			},
			in:  &JsonInput{parsedData: map[string]any{"name": "John", "lastName": "Doe", "age": 43}},
			out: &JsonInput{parsedData: map[string]any{"Name": "John", "Age": 43}},
		},
		{
			name: "Test missing key in input",
			p: &JsonToJsonProcessor{
				keyMappings: map[string]string{
					"name":     "Name",
					"lastName": "LastName",
					"age":      "Age",
				},
			},
			in:  &JsonInput{parsedData: map[string]any{"name": "John", "age": 43}},
			out: &JsonInput{parsedData: map[string]any{"Name": "John", "LastName": nil, "Age": 43}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Process(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Process() failed.\n  Got: %v.\n  Expected: %v", got, tt.out)
			}
		})
	}
}

func TestJsonToHttpProcessor(t *testing.T) {
	tests := []ProcessorTest{
		{
			name: "Test simple happy path",
			p: &JsonToHttpProcessor{
				method:  "GET",
				apiUrl:  "https://jsonplaceholder.typicode.com/posts/1",
				headers: map[string]string{},
			},
			in: &EmptyInput{},
			out: &HttpOutput{
				data: HttpResponseData{
					Ok:          true,
					StatusCode:  200,
					ContentType: "application/json; charset=utf-8",
					Body:        "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\n  \"body\": \"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"\n}",
				},
			},
		},
		{
			name: "Test simple happy path POST",
			p: &JsonToHttpProcessor{
				method:  "POST",
				apiUrl:  "https://jsonplaceholder.typicode.com/posts",
				headers: map[string]string{},
			},
			in: &JsonInput{
				parsedData: map[string]any{
					"userId": 1,
					"title":  "Do the dishes",
					"body":   "CLEEEAN!",
				},
			},
			out: &HttpOutput{
				data: HttpResponseData{
					Ok:          true,
					StatusCode:  201,
					ContentType: "application/json; charset=utf-8",
					Body:        "{\n  \"id\": 101\n}",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.Process(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Process() failed.\nGot: %v.\nExpected: %v", got, tt.out)
			}
		})
	}
}
