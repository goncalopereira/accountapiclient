package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

//Returns List of accounts based on parameters
//Noticed an empty list had data:null and thought about moving to empty array
//But according to https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
//Null array is preferred in Go
func (client *Client) List(parameters *url.Values) (data.IOutput, error) {
	requestURL, configErr := client.config.Accounts(parameters)
	if configErr != nil {
		return nil, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return nil, requestErr
	}
	if response.StatusCode == http.StatusOK {
		responseAccounts := &account.Accounts{}
		accountErr := json.BodyToData(response.Body, responseAccounts)
		if accountErr != nil {
			return nil, accountErr
		}

		return responseAccounts, nil
	}
	errorResponse := &data.ErrorResponse{}
	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return nil, errorResponseError
	}
	return errorResponse, nil
}