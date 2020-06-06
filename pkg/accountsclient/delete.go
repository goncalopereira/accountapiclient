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

	requestURL, err := client.config.Account(id, parameters)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest("DELETE", requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusNoContent {
		return &data.NoContent{}, nil
	}

	return errorResponseHandling(response)
}
