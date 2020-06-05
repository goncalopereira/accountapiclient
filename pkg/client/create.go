package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

func (client *Client) Create(accountRequest *account.Account) (data.IOutput, error) {
	requestData, dataErr := json.DataToBody(accountRequest)
	if dataErr != nil {
		return nil, dataErr
	}
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return nil, configErr
	}
	response, requestErr := client.Request.Post(requestURL.String(), requestData)
	if requestErr != nil {
		return nil, requestErr
	}

	if response.StatusCode == http.StatusCreated {
		responseAccount := &account.Account{}
		accountErr := json.BodyToData(response.Body, responseAccount)
		if accountErr != nil {
			return nil, accountErr
		}
		return responseAccount, nil
	}
	errorResponse := &data.ErrorResponse{}
	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return nil, errorResponseError
	}
	return errorResponse, nil
}
