package http

import (
	"io/ioutil"
	"net/http"
)

//build the request wrapper to be able to test net/http separately from service
//returns response data.
type Request struct {
	IRequest
	client *http.Client
}

func NewRequest() *Request {
	r := &Request{}
	r.client = &http.Client{}

	return r
}

func (h *Request) Get(req *http.Request) (*Response, error) {
	return h.handleRequest(req)
}

func (h *Request) handleRequest(req *http.Request) (*Response, error) {
	response, responseErr := h.client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}

	if responseErr != nil {
		return nil, responseErr
	}

	//did not test ReadAll error, would require mocking it
	responseBody, bodyErr := ioutil.ReadAll(response.Body)

	if responseErr != nil {
		return nil, bodyErr
	}

	return &Response{StatusCode: response.StatusCode, Body: responseBody}, nil
}
