package test

import (
	"net/http"

	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/mock"
)

//RequestMock to be able to mock net/http
//looked into using Pact but with only an external Provider it would be too much to force it to use e2e tests
//looked into using Sling but had to keep all third party libraries out of the client.
type RequestMock struct {
	mock.Mock
	internalhttp.IRequest
}

//NewRequestMock returns Request based on a given test Response.
//could improve by getting correct urls instead of any.
func NewRequestMock(response *internalhttp.Response, err error) internalhttp.IRequest {
	client := new(RequestMock)
	client.On("Do", mock.AnythingOfType("*http.Request")).Return(response, err).Once()

	return client
}

//Do should be mocked to wrap http.Client,
//and returns http.Response byte array already read in Response.
func (r *RequestMock) Do(req *http.Request) (*internalhttp.Response, error) {
	arguments := r.Called(req)
	errArgument := arguments.Get(1)

	var err error
	if errArgument != nil {
		err = errArgument.(error)
	}

	return arguments.Get(0).(*internalhttp.Response), err
}
