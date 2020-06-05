package http_test

import (
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/mock"
)

//need to be able to mock net/http
//looked into using Pact but with only an external Provider it would be too much to force it to use e2e tests
//looked into using Sling but had to keep all third party libraries out of the client.
type RequestMock struct {
	mock.Mock
	internalhttp.IRequest
}

//could improve by getting correct urls instead of any.
func NewGetRequestMock(response *internalhttp.Response, err error) internalhttp.IRequest {
	client := new(RequestMock)
	client.On("Get", mock.AnythingOfType("string")).Return(response, err).Once()

	return client
}
func NewDeleteRequestMock(response *internalhttp.Response, err error) internalhttp.IRequest {
	client := new(RequestMock)
	client.On("Delete", mock.AnythingOfType("string")).Return(response, err).Once()

	return client
}

func NewPostRequestMock(response *internalhttp.Response, err error) internalhttp.IRequest {
	client := new(RequestMock)
	client.On("Post", mock.AnythingOfType("string"), mock.Anything).Return(response, err).Once()

	return client
}

func (r *RequestMock) Get(endpoint string) (*internalhttp.Response, error) {
	arguments := r.Called(endpoint)
	return returnResponseAndError(arguments)
}

func (r *RequestMock) Post(endpoint string, requestData []byte) (*internalhttp.Response, error) {
	arguments := r.Called(endpoint, requestData)
	return returnResponseAndError(arguments)
}

func (r *RequestMock) Delete(endpoint string) (*internalhttp.Response, error) {
	arguments := r.Called(endpoint)
	return returnResponseAndError(arguments)
}

func returnResponseAndError(arguments mock.Arguments) (*internalhttp.Response, error) {
	errArgument := arguments.Get(1)

	var err error
	if errArgument != nil {
		err = errArgument.(error)
	}

	return arguments.Get(0).(*internalhttp.Response), err
}
