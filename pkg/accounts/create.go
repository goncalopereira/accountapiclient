package accounts

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

func (client *Client) Create(accountRequest *account.Data) (data.IOutput, error) {
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return &data.NoOp{}, configErr
	}

	requestData, dataErr := json.DataToBody(accountRequest)
	if dataErr != nil {
		return nil, dataErr
	}

	response, responseErr := client.handleRequest("POST", requestURL.String(), requestData)
	if responseErr != nil {
		return &data.NoOp{}, responseErr
	}

	if response.StatusCode == http.StatusCreated {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
