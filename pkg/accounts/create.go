package accounts

import (
	"bytes"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
)

func (client *Client) Create(accountRequest *account.Data) (data.IOutput, error) {
	requestData, dataErr := json.DataToBody(accountRequest)
	if dataErr != nil {
		return nil, dataErr
	}

	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return nil, configErr
	}

	req, err := http.NewRequest("POST", requestURL.String(), bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	response, requestErr := client.Request.Do(req)
	if requestErr != nil {
		return nil, requestErr
	}

	if response.StatusCode == http.StatusCreated {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
