package main

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

// Action is our concrete type that will become the adapter for our HTTP handlers
type Action func(r *http.Request) *Response

func (handler Action) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := handler(r)

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

// Error decorator
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

func ErrorJSON(status int, err error, headers Headers) *Response {
	errResp := struct {
		Error error `json:"error"`
	}{
		Error: err,
	}

	b, err := json.Marshal(errResp)

	if err != nil {
		return Error(http.StatusInternalServerError, err, headers)
	}

	return &Response{
		Status:      status,
		ContentType: "application/json",
		Content:     bytes.NewBuffer(b),
		Headers:     headers,
	}
}

func DataJSON(status int, v interface{}, headers Headers) *Response {
	b, err := json.Marshal(v)

	if err != nil {
		return ErrorJSON(http.StatusInternalServerError, err, headers)
	}

	return Data(status, b, headers)
}
