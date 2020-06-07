package test

import (
	"github.com/goncalopereira/accountapiclient/internal/api"
	"github.com/goncalopereira/accountapiclient/internal/http"
	accountsClient "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
)

//NewTestClient is used to inject a mocked Request during tests.
func NewTestClient(request http.IRequest) *accountsClient.Client {
	return &accountsClient.Client{Config: api.DefaultAPI(), Request: request}
}
