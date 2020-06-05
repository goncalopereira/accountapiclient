package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"net/url"
)

func (client *Client) List(parameters *url.Values) (*data.Output, error) {
	requestURL, configErr := client.config.Accounts(parameters)
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}
