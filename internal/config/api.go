package config

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/env"
	"net/url"
)
type API struct {
	host   string
	port   string
	scheme string
	accountsURL string
	healthURL string
}


func NewAPI(host string, port string, scheme string) *API {
	a := &API{host: host, port: port, scheme: scheme}
	a.accountsURL = "/v1/organisation/accounts"
	a.healthURL = "/v1/health1"
	return a
}

func DefaultAPI() *API {
	return NewAPI(env.GetEnv("API_HOST", "localhost"), env.GetEnv("API_PORT", "8080"), env.GetEnv("API_SCHEME", "http"))
}

func (c *API) BaseURL() string {
	return c.scheme + "://" + c.host + ":" + c.port
}

func (c *API) Accounts(parameters *url.Values) (*url.URL, error) {
	requestUrl := fmt.Sprintf("%s%s", c.BaseURL(), c.accountsURL)
	return BuildUrl(requestUrl, parameters)

}

func (c *API) Account(id string, parameters *url.Values) (*url.URL, error) {
	requestUrl := fmt.Sprintf("%s%s/%s", c.BaseURL(), c.accountsURL, id)
	return BuildUrl(requestUrl, parameters)
}

func BuildUrl(requestUrl string, parameters *url.Values) (*url.URL, error) {
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}

	u.RawQuery = parameters.Encode()

	return u, nil
}