package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

func (client *Client) Create(accountRequest *account.Account) (*data.Output, error) {
	requestData, dataErr := json.DataToBody(accountRequest)
	if dataErr != nil {
		return &data.Output{}, dataErr
	}
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Post(requestURL.String(), requestData)
	if requestErr != nil {
		return &data.Output{}, requestErr
	}

	output := &data.Output{}
	if response.StatusCode == http.StatusCreated {
		responseAccount := &account.Account{}
		accountErr := json.BodyToData(response.Body, responseAccount)
		if accountErr != nil {
			return output, accountErr
		}
		output.Account = responseAccount
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
