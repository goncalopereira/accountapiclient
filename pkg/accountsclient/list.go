package accountsclient

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
	"net/url"
)

//Returns List of accounts based on parameters
//Noticed an empty list had data:null and thought about moving to empty array
//But according to https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
//Null array is preferred in Go
//FakeAPI does not respect filters only pagination, unit tests include filter not e2e
//Did not add a parameter filter to pick the available ones as the API might change this quickly
//There is the issue of developers mistyping filter names but that's a testing issue.
func (client *Client) List(parameters *url.Values) (data.IOutput, error) {
	requestURL, err := client.config.Accounts(parameters)
	if err != nil {
		return &data.NoOp{}, err
	}

	response, err := client.handleRequest("GET", requestURL.String(), nil)
	if err != nil {
		return &data.NoOp{}, err
	}

	if response.StatusCode == http.StatusOK {
		return validResponseHandling(response, &account.AccountsData{})
	}

	return errorResponseHandling(response)
}
