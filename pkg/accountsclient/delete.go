package accountsclient

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/google/uuid"
)

//Delete deletes an Account based on ID and Version,
//returns IOutput with Deleted, or NoOp when error.
func (client *Client) Delete(id uuid.UUID, version int) (data.IOutput, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))

	requestURL, err := client.Config.Account(id, parameters)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodDelete, requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusNoContent {
		return &data.Deleted{}, nil
	}

	return errorResponseHandling(response)
}
