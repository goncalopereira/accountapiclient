package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//all use of net/http
//external http request for endpoint
//returns response data
type Request struct {
	client *http.Client
}

func NewRequest() *Request {
	r := &Request{}
	r.client = &http.Client{}
	return r
}

func (h *Request) Get(endpoint string) (*Response, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return h.handleRequest(endpoint, req)
}

func (h *Request) Delete(endpoint string) (*Response, error) {
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return nil, err
	}
	return h.handleRequest(endpoint, req)
}

func (h *Request) Post(endpoint string, requestData []byte) (*Response, error) {
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}
	return h.handleRequest(endpoint, req)
}

func (h *Request) handleRequest(endpoint string, req *http.Request) (*Response, error) {
	response, responseErr := h.client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}

	if responseErr != nil {
		return NewBadResponse(), responseErr
	}

	//did not test ReadAll error, would require mocking it
	responseBody, bodyErr := ioutil.ReadAll(response.Body)

	if responseErr != nil {
		return NewBadResponse(), fmt.Errorf("request %v body: %v", endpoint, bodyErr.Error())
	}

	return NewResponse(response.StatusCode, responseBody), nil
}
