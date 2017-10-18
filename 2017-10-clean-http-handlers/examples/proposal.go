package main

import (
	"bytes"
	"io"
	"net/http"
)

// START 1 OMIT
func SimpleHandler(r *http.Request) *Response {
	content := bytes.NewBufferString("{}")
	return &Response{200, "application/json", content, nil}
}

// Headers is a map of string to string where the key is the the header name
// and the value is its value.
type Headers map[string]string

// Response is the generic response object for our handlers
type Response struct {
	Status      int       // StatusCode
	ContentType string    // Content Type to writer
	Content     io.Reader // Content to be written to the response writer
	Headers     Headers   // Headers to be written to the response writer
}

// END 1 OMIT
