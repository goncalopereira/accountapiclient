package client

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
	response, requestErr := client.Request.Delete(requestURL.String())
	if requestErr != nil {
		return nil, requestErr
	}
	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	return errorResponseHandling(response)
}
