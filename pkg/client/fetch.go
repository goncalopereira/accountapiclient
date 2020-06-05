package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"log"
	"net/http"
	"net/url"
)

func (client *Client) Fetch(id string) (*data.Output, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return &data.Output{}, configErr
	}
	log.Print(requestURL.String())
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}

	output := &data.Output{}
	if response.StatusCode == http.StatusOK {
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
