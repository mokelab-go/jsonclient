package mock

import (
	j "github.com/mokelab-go/jsonclient"
)

type MockClient struct {
	Url     string
	Mehtod  string
	Headers map[string]string
	Body    map[string]string

	respStatus  int
	respHeaders map[string]string
	respBody    map[string]string
	respError   error
}

func NewClient() *Mockclient {
	return &MockClient{
		Headers: make(map[string]string),
	}
}

func (c *MockClient) SetURL(url string) {
	c.Url = url
}

func (c *MockClient) SetMethod(method string) {
	c.Mehtod = method
}

func (c *MockClient) SetHeader(key, value string) {
	c.Headers[key] = value
}

func (c *MockClient) Send(body map[string]interface{}) (j.Response, error) {
	c.Body = body
	return j.Response{
		Status:  c.respStatus,
		Headers: c.respHeaders,
		Body:    c.respBody,
	}, c.respError
}

func (c *MockClient) SetResponse(status int, headers, body map[string]string) {
	c.respStatus = status
	c.respHeaders = headers
	c.respBody = body
	c.respError = nil
}

func (c *MockClient) SetErrorResponse(err error) {
	c.respError = err
}
