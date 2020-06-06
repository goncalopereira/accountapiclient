package accountsclient

import (
	"bytes"
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"net/http"
)

type Client struct {
	config  config.IAPI
	Request internalhttp.IRequest
}

func NewClient(config config.IAPI, request internalhttp.IRequest) *Client {
	return &Client{config: config, Request: request}
}

func errorResponseHandling(response *internalhttp.Response) (data.IOutput, error) {
	errorResponse := &data.ErrorResponse{StatusCode: response.StatusCode}

	err := json.Unmarshal(response.Body, &errorResponse)
	if err != nil {
		return &data.NoOp{}, err
	}

	return errorResponse, nil
}

func validResponseHandling(response *internalhttp.Response, responseData data.IOutput) (data.IOutput, error) {
	err := json.Unmarshal(response.Body, responseData)
	if err != nil {
		return &data.NoOp{}, err
	}

	return responseData, nil
}

func (client *Client) handleRequest(
	method string,
	requestURL string,
	data []byte) (*internalhttp.Response, error) {
	req, err := http.NewRequest(method, requestURL, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	response, err := client.Request.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}
