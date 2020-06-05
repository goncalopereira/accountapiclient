package config_test

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/stretchr/testify/mock"
	"net/url"
)

//API mock required to wrap around environment config issues
type APIMock struct {
	mock.Mock
	config.IAPI
}

//ideally have a better implementation than just mock anything
func NewAPIMock(returnURL *url.URL, returnError error) config.IAPI {
	api := new(APIMock)
	api.On("Accounts", mock.AnythingOfType("*url.Values")).
		Return(returnURL, returnError).Once()
	api.On("Account", mock.AnythingOfType("string"), mock.AnythingOfType("*url.Values")).
		Return(returnURL, returnError).Once()

	return api
}

func (a *APIMock) Accounts(parameters *url.Values) (*url.URL, error) {
	arguments := a.Called(parameters)
	return returnResponseAndError(arguments)
}
func (a *APIMock) Account(id string, parameters *url.Values) (*url.URL, error) {
	arguments := a.Called(id, parameters)
	return returnResponseAndError(arguments)
}

func returnResponseAndError(arguments mock.Arguments) (*url.URL, error) {
	errArgument := arguments.Get(1)

	var err error
	if errArgument != nil {
		err = errArgument.(error)
	}

	return arguments.Get(0).(*url.URL), err
}
