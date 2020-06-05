package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/url"
)

func (client *Client) Create(account *account.Account) (*data.Output, error) {
	requestData, dataErr := json.DataToBody(account)
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
	return handleResponse(response)
}
