package accounts

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

func (client *Client) Fetch(id string) (data.IOutput, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return nil, configErr
	}

	response, responseErr := client.handleRequest("GET", requestURL, nil)
	if responseErr != nil {
		return nil, responseErr
	}
	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
