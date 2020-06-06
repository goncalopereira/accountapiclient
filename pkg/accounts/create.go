package accounts

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

func (client *Client) Create(accountRequest *account.Data) (data.IOutput, error) {
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return &data.NoOp{}, configErr
	}

	response, responseErr := client.handleRequest("POST", requestURL.String(), accountRequest)
	if responseErr != nil {
		return &data.NoOp{}, responseErr
	}

	if response.StatusCode == http.StatusCreated {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
