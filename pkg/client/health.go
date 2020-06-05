package client

import "github.com/goncalopereira/accountapiclient/internal/data"

func (client *Client) Health() (*data.Output, error) {
	requestURL, configErr := client.config.Health()
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}
