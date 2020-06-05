package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

func (client *Client) Fetch(id string) (data.IOutput, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return nil, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return nil, requestErr
	}

	if response.StatusCode == http.StatusOK {
		responseAccount := &account.Data{}
		accountErr := json.BodyToData(response.Body, responseAccount)
		if accountErr != nil {
			return nil, accountErr
		}
		return responseAccount, nil
	}
	return errorResponseHandling(response)
}
