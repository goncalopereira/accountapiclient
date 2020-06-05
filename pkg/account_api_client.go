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

type Output struct {
	Account       *account.Account
	ErrorResponse *data.ErrorResponse
	Accounts      *[]account.Account
}

//default new client with real http
func NewAccountAPIClient() *Client {
	client := &Client{config: config.DefaultAPI()}
	client.request = http.NewRequest()
	return client
}

func (client *Client) Fetch(id string) (*Output, error) {
	requestURL, configErr := client.config.Account(id, &url.Values{})
	if configErr != nil {
		return &Output{}, configErr
	}
	log.Print(requestURL.String())
	response, requestErr := client.request.Get(requestURL)
	if requestErr != nil {
		return &Output{}, requestErr
	}
	return handleResponse(response)
}

func handleResponse(response *http.Response) (*Output, error) {
	log.Print(strconv.Itoa(response.StatusCode))
	log.Print(string(response.Body))
	return &Output{}, nil
}

func (client *Client) Create(account *account.Account) (*Output, error) {
	requestData, dataErr := json.DataToBody(account)
	if dataErr != nil {
		return &Output{}, dataErr
	}
	requestURL, configErr := client.config.Accounts(&url.Values{})
	if configErr != nil {
		return &Output{}, configErr
	}
	response, requestErr := client.request.Post(requestURL, requestData)
	if requestErr != nil {
		return &Output{}, requestErr
	}
	return handleResponse(response)
}

func (client *Client) List(parameters *url.Values) (*Output, error) {
	requestURL, configErr := client.config.Accounts(parameters)
	if configErr != nil {
		return &Output{}, configErr
	}
	response, requestErr := client.request.Get(requestURL)
	if requestErr != nil {
		return &Output{}, requestErr
	}
	return handleResponse(response)
}
func (client *Client) Delete(id string, version int) (*Output, error) {
	parameters := &url.Values{}
	parameters.Add("version", strconv.Itoa(version))
	requestURL, configErr := client.config.Account(id, parameters)
	if configErr != nil {
		return &Output{}, configErr
	}
	response, requestErr := client.request.Delete(requestURL)
	if requestErr != nil {
		return &Output{}, requestErr
	}
	return handleResponse(response)
}

func (client *Client) Health() (*Output, error) {
	requestURL, configErr := client.config.Health()
	if configErr != nil {
		return &Output{}, configErr
	}
	response, requestErr := client.request.Get(requestURL)
	if requestErr != nil {
		return &Output{}, requestErr
	}
	return handleResponse(response)
}
