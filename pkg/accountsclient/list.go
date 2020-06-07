package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"net/http"
	"net/url"
)

//List returns a list of accounts based on the parameters (filters),
//returns Accounts, ErrorMessage or NoOp when error.
//DEV COMMENTS:
//Noticed an empty list had data:null and thought about moving to empty array
//but according to https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
//nil array is preferred in Go
//FakeAPI does not respect filters only pagination, unit tests include filter not e2e,
//did not add a parameter filter to pick the available ones as the API might change this quickly
//there is the issue of developers mistyping filter names but that's a testing issue.
func (client *Client) List(parameters *url.Values) (data.IOutput, error) {
	requestURL, err := client.Config.Accounts(parameters)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &data.AccountsData{})
	}

	return errorResponseHandling(response)
}
