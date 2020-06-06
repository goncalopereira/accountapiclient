package accounts

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"net/http"
	"net/url"
	"strconv"
)

func (client *Client) Delete(id string, version int) (data.IOutput, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))

	requestURL, configErr := client.config.Account(id, parameters)
	if configErr != nil {
		return nil, configErr
	}

	response, responseErr := client.handleRequest("GET", requestURL, nil)
	if responseErr != nil {
		return nil, responseErr
	}

	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	return errorResponseHandling(response)
}
