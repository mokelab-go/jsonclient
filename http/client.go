package http

import (
	j "../"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type client struct {
	url     string
	mehtod  string
	headers map[string]string
}

func NewClient() *client {
	return &client{
		headers: make(map[string]string),
	}
}

func (c *client) SetURL(url string) {
	c.url = url
}

func (c *client) SetMethod(method string) {
	c.mehtod = method
}

func (c *client) SetHeader(key, value string) {
	c.headers[key] = value
}

func (c *client) Send(body map[string]interface{}) (j.Response, error) {
	var r io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return j.Response{}, err
		}
		r = bytes.NewReader(b)
	}
	req, err := http.NewRequest(c.mehtod, c.url, r)
	if err != nil {
		return j.Response{}, err
	}
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
	// send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return j.Response{}, err
	}
	doc := json.NewDecoder(resp.Body)
	var obj map[string]interface{}
	err = doc.Decode(&obj)
	if err != nil {
		return j.Response{}, err
	}
	headers := make(map[string]string)
	for key, _ := range resp.Header {
		headers[key] = resp.Header.Get(key)
	}
	return j.Response{
		Status:  resp.StatusCode,
		Headers: headers,
		Body:    obj,
	}, nil
}
