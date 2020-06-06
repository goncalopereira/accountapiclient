package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

func (client *Client) Fetch(id string) (data.IOutput, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return &data.NoOp{}, configErr
	}

	response, responseErr := client.handleRequest("GET", requestURL.String(), nil)
	if responseErr != nil {
		return &data.NoOp{}, responseErr
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
