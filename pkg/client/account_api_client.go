package client

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"log"
	"net/http"
	"net/url"
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
func NewDefaultAccountAPIClient() *Client {
	return NewClient(config.DefaultAPI(), internalhttp.NewRequest())
}

func (client *Client) Fetch(id string) (*data.Output, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return &data.Output{}, configErr
	}
	log.Print(requestURL.String())
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}

	output := &data.Output{}
	if response.StatusCode == http.StatusOK {
		responseAccount := &account.Account{}
		accountErr := json.BodyToData(response.Body, responseAccount)
		if accountErr != nil {
			return output, accountErr
		}
		output.Account = responseAccount
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

func (client *Client) Create(account *account.Account) (*data.Output, error) {
	requestData, dataErr := json.DataToBody(account)
	if dataErr != nil {
		return &data.Output{}, dataErr
	}
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Post(requestURL.String(), requestData)
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}

func (client *Client) List(parameters *url.Values) (*data.Output, error) {
	requestURL, configErr := client.config.Accounts(parameters)
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.Request.Get(requestURL.String())
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}
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

func handleResponse(response *internalhttp.Response) (*data.Output, error) {
	log.Print(strconv.Itoa(response.StatusCode))
	log.Print(string(response.Body))

	return &data.Output{}, nil
}
