package flow

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type response struct {
	status  int
	headers map[string]string
	body    string
}

var (
	responseChannel = make(chan response, 1)
	requests        int
)

func testServerHandler(w http.ResponseWriter, r *http.Request) {
	requests += 1

	res := <-responseChannel

	for k, v := range res.headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(res.status)

	w.Write([]byte(res.body))
}

func TestHttpToJsonProcessor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(testServerHandler))
	defer server.Close()

	tests := []struct {
		name string
		res  response
		p    Processor
		in   Input
		out  Output
	}{
		{
			name: "Test happy path",
			res: response{
				status: 200,
				headers: map[string]string{
					"Content-Type": "application/json",
				},
				body: `{"msg":"Hello world!"}`,
			},
			p: &HttpToJsonProcessor{
				method:  "GET",
				apiUrl:  server.URL,
				headers: map[string]string{},
			},
			in: nil,
			out: &JsonInput{
				rawData:    []byte("{\"msg\":\"Hello world!\"}"),
				parsedData: map[string]any{"msg": "Hello world!"},
			},
		},
	}

	for _, tt := range tests {
		requests = 0
		t.Run(tt.name, func(t *testing.T) {
			responseChannel <- tt.res
			got := tt.p.Process(tt.in)
			if got.HasError() {
				t.Fatalf("Process() failed. Output had error: %v", got.Error())
			}

			if requests != 1 {
				t.Errorf("Process() = More than 1 request was made to the server")
			}

			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("Process() failed.\nGot: %v.\nExpected: %v", got, tt.out)
			}
		})
	}
}
