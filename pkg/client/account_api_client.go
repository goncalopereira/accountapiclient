package client

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"log"
	"strconv"
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

func handleResponse(response *internalhttp.Response) (*data.Output, error) {
	log.Print(strconv.Itoa(response.StatusCode))
	log.Print(string(response.Body))

	return &data.Output{}, nil
}
