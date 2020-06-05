package client

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
)

type Client struct {
	config  config.IAPI
	Request internalhttp.IRequest
}

func NewClient(config config.IAPI, request internalhttp.IRequest) *Client {
	return &Client{config: config, Request: request}
}

//default new client with real http
func NewClientFromEnv() *Client {
	return NewClient(config.DefaultAPI(), internalhttp.NewRequest())
}
