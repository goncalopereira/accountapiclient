package accounts

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
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

	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return nil, errorResponseError
	}

	return errorResponse, nil
}

func validResponseHandling(response *internalhttp.Response, responseData data.IOutput) (data.IOutput, error) {
	accountErr := json.BodyToData(response.Body, responseData)
	if accountErr != nil {
		return nil, accountErr
	}

	return responseData, nil
}
