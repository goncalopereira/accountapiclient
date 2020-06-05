package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
	"strconv"
)

func (client *Client) Delete(id string, version int) (*data.ErrorResponse, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))
	requestURL, configErr := client.config.Account(id, parameters)
	if configErr != nil {
		return nil, configErr
	}
	response, requestErr := client.Request.Delete(requestURL.String())
	if requestErr != nil {
		return nil, requestErr
	}
	if response.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	errorResponse := &data.ErrorResponse{StatusCode: response.StatusCode}
	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return nil, errorResponseError
	}
	return errorResponse, nil
}
