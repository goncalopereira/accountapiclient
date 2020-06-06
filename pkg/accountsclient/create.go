package accountsclient

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

//Create receives the Data type with Account
//Returns IOutput with Account ErrorMessage, or NoOp when error.
func (client *Client) Create(accountRequest *account.Data) (data.IOutput, error) {
	requestURL, err := client.config.Accounts(&url.Values{})
	if err != nil {
		return &data.NoOp{}, err
	}

	requestData, err := json.Marshal(accountRequest)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodPost, requestURL.String(), requestData)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusCreated {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
