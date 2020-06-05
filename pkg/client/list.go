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
func (client *Client) List(parameters *url.Values) (*data.Output, error) {
	requestURL, configErr := client.config.Accounts(parameters)
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	output := &data.Output{}
	if response.StatusCode == http.StatusOK {
		responseAccounts := &account.Accounts{}
		accountErr := json.BodyToData(response.Body, responseAccounts)
		if accountErr != nil {
			return output, accountErr
		}

		output.Accounts = responseAccounts
		return output, nil
	}
	errorResponse := &data.ErrorResponse{}
	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return output, errorResponseError
	}
	output.ErrorResponse = errorResponse
	return output, nil
}
