package test

import (
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/mock"
	"net/http"
)

//need to be able to mock net/http
//looked into using Pact but with only an external Provider it would be too much to force it to use e2e tests
//looked into using Sling but had to keep all third party libraries out of the client
type RequestMock struct {
	mock.Mock
	internalhttp.IRequest
	Response internalhttp.Response
	Err      error
}

func NewGetRequestMock(response internalhttp.Response, err error) *RequestMock {
	client := new(RequestMock)
	client.On("Get", "http://localhost:8080/endpoint").Return(http.Response{}, nil)
	client.Response = response
	client.Err = err
	return client
}

func NewDeleteRequestMock(response internalhttp.Response, err error) *RequestMock {
	client := new(RequestMock)
	client.On("Delete", "http://localhost:8080/endpoint").Return(http.Response{}, nil)
	client.Response = response
	client.Err = err
	return client
}

func NewPostRequestMock(requestData []byte, response internalhttp.Response, err error) *RequestMock {
	client := new(RequestMock)
	client.On("Post", "http://localhost:8080/endpoint", requestData).Return(http.Response{}, nil)
	client.Response = response
	client.Err = err
	return client
}

func (r *RequestMock) Get(endpoint string) (*internalhttp.Response, error) {
	_ = r.Called(endpoint)
	return &r.Response, r.Err
}

func (r *RequestMock) Post(endpoint string, requestData []byte) (*internalhttp.Response, error) {
	_ = r.Called(endpoint, requestData)
	return &r.Response, r.Err
}

func (r *RequestMock) Delete(endpoint string) (*internalhttp.Response, error) {
	_ = r.Called(endpoint, endpoint)
	return &r.Response, r.Err
}
