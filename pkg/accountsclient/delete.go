package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"net/http"
	"net/url"
	"strconv"
)

//Delete receives the account id and version
//Returns IOutput with NoContent, or NoOp when error.
func (client *Client) Delete(id string, version int) (data.IOutput, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))

	requestURL, err := client.config.Account(id, parameters)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusNoContent {
		return &data.NoContent{}, nil
	}

	return errorResponseHandling(response)
}
