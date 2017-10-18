package adapter

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Headers is a map of string to string where the key is the the header name
// and the value is its value.
type Headers map[string]string

// Response is the generic response object for our handlers
type Response struct {
	// StatusCode
	Status int
	// Content Type to writer
	ContentType string
	// Content to be written to the response writer
	Content io.Reader
	// Headers to be written to the response writer
	Headers Headers
}

// START 1 OMIT
type Action func(r *http.Request) *Response

// END 1 OMIT

// START 2 OMIT
func (handler Action) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := handler(r) // Call our handler and get the Response

	// No response given, answer with 204 No Content
	if response == nil {
		rw.WriteHeader(http.StatusNoContent)
		return
	}

	if response.ContentType != "" {
		rw.Header().Set("Content-Type", response.ContentType)
	}

	for k, v := range response.Headers {
		rw.Header().Set(k, v)
	}

	rw.WriteHeader(response.Status)

	_, err := io.Copy(rw, response.Content)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// END 2 OMIT

// START 3 OMIT
func Data(status int, content []byte, headers Headers) *Response {
	return &Response{
		Status:  status,
		Content: bytes.NewBuffer(content),
		Headers: headers,
	}
}

func Error(status int, err error, headers Headers) *Response {
	return &Response{
		Status:  status,
		Content: bytes.NewBufferString(err.Error()),
		Headers: headers,
	}
}

// END 3 OMIT

// START 4 OMIT
func DataJSON(status int, v interface{}, headers Headers) *Response {
	b, err := json.Marshal(v)
	if err != nil {
		return ErrorJSON(http.StatusInternalServerError, err, headers)
	}

	return &Response{
		Status:      status,
		ContentType: "application/json",
		Content:     bytes.NewBuffer(b),
		Headers:     headers,
	}
}

type jsonErr struct {
	Error error `json:"error"`
}

func ErrorJSON(status int, err error, headers Headers) *Response {
	resp := jsonErr{Error: err}

	return DataJSON(http.StatusInternalServerError, resp, nil)
}

// END 4 OMIT
