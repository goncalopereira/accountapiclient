package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

//Fetch receives the Account id
//Returns IOutput with Account, ErrorMessage, or NoOp when error.
func (client *Client) Fetch(id string) (data.IOutput, error) {
	requestURL, err := client.config.Account(id, &url.Values{})
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
