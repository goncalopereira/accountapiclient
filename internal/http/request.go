package http

import (
	"io/ioutil"
	"net/http"
)

//Request builds the request wrapper to be able to test net/http separately from service
//returns Response data.
type Request struct {
	IRequest
	*http.Client
}

//NewClient builds the Request type by default with real http.Client.
func NewClient() *Request {
	r := &Request{}
	r.Client = &http.Client{}

	return r
}

//Do wraps the net.http.client Do method to be mock-able
//also handles the response reading to be able to close http.response.body
//returns a plain response type with just status code and body byte array.
func (h *Request) Do(req *http.Request) (*Response, error) {
	response, err := h.Client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	//did not test ReadAll error, would require mocking it
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return &Response{StatusCode: response.StatusCode, Body: responseBody}, nil
}
