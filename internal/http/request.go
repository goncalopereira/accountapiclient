package http

import (
	"io/ioutil"
	"net/http"
)

//build the request wrapper to be able to test net/http separately from service
//returns response data.
type Request struct {
	IRequest
	*http.Client
}

func NewRequest() *Request {
	r := &Request{}
	r.Client = &http.Client{}

	return r
}

//wraps the Do method to be mockabble
//also handles the response reading to be able to close http.response.body
//returns a plain response type with just status code and body byte array.
func (h *Request) Do(req *http.Request) (*Response, error) {
	response, responseErr := h.Client.Do(req)
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
