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

	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	response, requestErr := client.Request.Do(req)
	if requestErr != nil {
		return nil, requestErr
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &account.Data{})
	}

	return errorResponseHandling(response)
}
