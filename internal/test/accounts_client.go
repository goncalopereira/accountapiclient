package test

import (
	"github.com/goncalopereira/accountapiclient/internal/api"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	accountsClient "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
)

func NewTestClient(request internalhttp.IRequest) *accountsClient.Client {
	return &accountsClient.Client{Config: api.DefaultAPI(), Request: request}
}
