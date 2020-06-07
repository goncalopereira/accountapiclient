package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

//Fetch returns an Account based on ID,
//returns IOutput with Account, ErrorMessage, or NoOp when error.
func (client *Client) Fetch(id uuid.UUID) (data.IOutput, error) {
	requestURL, err := client.Config.Account(id, &url.Values{})
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &data.Data{})
	}

	return errorResponseHandling(response)
}
