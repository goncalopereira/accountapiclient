package config

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/env"
	"net/url"
)

type API struct {
	IAPI
	host        string
	port        string
	scheme      string
	accountsURL string
	healthURL   string
}

func NewAPI(host string, port string, scheme string) IAPI {
	a := &API{host: host, port: port, scheme: scheme}
	a.accountsURL = "/v1/organisation/accounts"
	a.healthURL = "/v1/health1"
	return a
}

func DefaultAPI() IAPI {
	return NewAPI(env.GetEnv("API_HOST", "localhost"),
		env.GetEnv("API_PORT", "8080"),
		env.GetEnv("API_SCHEME", "http"))
}

func (c *API) baseURL() string {
	return c.scheme + "://" + c.host + ":" + c.port
}

func (c *API) Accounts(parameters *url.Values) (*url.URL, error) {
	requestURL := fmt.Sprintf("%s%s", c.baseURL(), c.accountsURL)
	return buildURL(requestURL, parameters)
}

func (c *API) Account(id string, parameters *url.Values) (*url.URL, error) {
	requestURL := fmt.Sprintf("%s%s/%s", c.baseURL(), c.accountsURL, id)
	return buildURL(requestURL, parameters)
}

func (c *API) Health() (*url.URL, error) {
	requestURL := fmt.Sprintf("%s%s", c.baseURL(), c.healthURL)
	return buildURL(requestURL, &url.Values{})
}

func buildURL(requestURL string, parameters *url.Values) (*url.URL, error) {
	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	u.RawQuery = parameters.Encode()

	return u, nil
}
