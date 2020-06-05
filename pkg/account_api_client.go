package accountapiclient

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"log"
	"net/url"
	"strconv"
)

type Client struct {
	config  *config.API
	request http.IRequest
}

func NewClient(config *config.API, request http.IRequest) *Client {
	return &Client{config: config, request: request}
}

//default new client with real http
func NewDefaultAccountAPIClient() *Client {
	return NewClient(config.DefaultAPI(), http.NewRequest())
}

func (client *Client) Fetch(id string) (*data.Output, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return &data.Output{}, configErr
	}
	log.Print(requestURL.String())
	response, requestErr := client.request.Get(requestURL)
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}

func handleResponse(response *http.Response) (*data.Output, error) {
	log.Print(strconv.Itoa(response.StatusCode))
	log.Print(string(response.Body))
	return &data.Output{}, nil
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
	response, requestErr := client.request.Post(requestURL, requestData)
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
	response, requestErr := client.request.Get(requestURL)
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
	response, requestErr := client.request.Delete(requestURL)
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}

func (client *Client) Health() (*data.Output, error) {
	requestURL, configErr := client.config.Health()
	if configErr != nil {
		return &data.Output{}, configErr
	}
	response, requestErr := client.request.Get(requestURL)
	if requestErr != nil {
		return &data.Output{}, requestErr
	}
	return handleResponse(response)
}
