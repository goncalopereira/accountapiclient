package http_test

import (
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/mock"
	"net/http"
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
	client.On("Get", mock.AnythingOfType("*http.Request")).Return(response, err).Once()

	return client
}

func NewPostRequestMock(response *internalhttp.Response, err error) internalhttp.IRequest {
	client := new(RequestMock)
	client.On("Post", mock.AnythingOfType("string"), mock.Anything).Return(response, err).Once()

	return client
}

func (r *RequestMock) Get(req *http.Request) (*internalhttp.Response, error) {
	arguments := r.Called(req)
	return returnResponseAndError(arguments)
}

func (r *RequestMock) Post(endpoint string, requestData []byte) (*internalhttp.Response, error) {
	arguments := r.Called(endpoint, requestData)
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
