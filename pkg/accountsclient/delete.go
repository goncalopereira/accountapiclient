package accountsclient

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
		return &data.NoOp{}, configErr
	}

	response, responseErr := client.handleRequest("DELETE", requestURL.String(), nil)
	if responseErr != nil {
		return &data.NoOp{}, responseErr
	}

	if response.StatusCode == http.StatusNoContent {
		return &data.NoContent{}, nil
	}

	return errorResponseHandling(response)
}
