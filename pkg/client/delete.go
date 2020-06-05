package client

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"net/http"
	"net/url"
	"strconv"
)

func (client *Client) Delete(id string, version int) (*data.Output, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))
	requestURL, configErr := client.config.Account(id, parameters)
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Delete(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	output := &data.Output{}
	if response.StatusCode == http.StatusNoContent {
		return output, nil
	}
	errorResponse := &data.ErrorResponse{}
	errorResponseError := json.BodyToData(response.Body, errorResponse)
	if errorResponseError != nil {
		return output, errorResponseError
	}
	output.ErrorResponse = errorResponse
	return output, nil
}
