//nolint:gosec
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
}

func NewRequest() *Request {
	return &Request{}
}

func (h *Request) Get(endpoint string) (*Response, error) {
	response, responseErr := http.Get(endpoint)

	return handleResponse(endpoint, response, responseErr)
}

func (h *Request) Post(endpoint string, requestData []byte) (*Response, error) {
	response, responseErr := http.Post(endpoint, "application/vnd.api+json", bytes.NewBuffer(requestData))

	return handleResponse(endpoint, response, responseErr)
}

func handleResponse(endpoint string, response *http.Response, responseErr error) (*Response, error) {
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
