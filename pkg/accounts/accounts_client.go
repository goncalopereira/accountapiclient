package accounts

import (
	"bytes"
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
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

	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return &data.NoOp{}, errorResponseError
	}

	return errorResponse, nil
}

func validResponseHandling(response *internalhttp.Response, responseData data.IOutput) (data.IOutput, error) {
	accountErr := json.BodyToData(response.Body, responseData)
	if accountErr != nil {
		return &data.NoOp{}, accountErr
	}

	return responseData, nil
}

func (client *Client) handleRequest(
	method string,
	requestURL string,
	data interface{}) (*internalhttp.Response, error) {
	requestData, dataErr := json.DataToBody(data)
	if dataErr != nil {
		return nil, dataErr
	}

	req, requestErr := http.NewRequest(method, requestURL, bytes.NewBuffer(requestData))
	if requestErr != nil {
		return nil, requestErr
	}

	response, responseErr := client.Request.Do(req)
	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}
