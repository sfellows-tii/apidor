package http

import (
	"bytes"
	"net/http"
	"time"
)

// RequestOptions is a http struct that holds the data that builds the HTTP request
type RequestOptions struct {
	Method        string
	BaseURI       string
	Endpoint      string
	ContentType   string
	Headers       map[string]string
	RequestParams map[string]string
	BodyParams    map[string]interface{}
	IsProxy       bool
	ProxyURI      string
}

// CreateRequest is a http function for creating a HTTP request based on request options
func CreateRequest(requestOptions RequestOptions) (*http.Request, error) {
	uri := buildURI(requestOptions.BaseURI, requestOptions.Endpoint, requestOptions.RequestParams)
	body, err := buildBody(requestOptions.ContentType, requestOptions.BodyParams)
	if err != nil {
		return nil, err
	}
	contentType, err := buildContentType(requestOptions.ContentType)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(requestOptions.Method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for headerName, headerValue := range requestOptions.Headers {
		req.Header.Set(headerName, headerValue)
	}
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

// SendRequest is a http function for sending a HTTP request and returns the response
func SendRequest(req *http.Request) (*http.Response, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
