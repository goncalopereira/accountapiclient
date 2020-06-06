package config

import (
	"fmt"
	"net/url"
	"os"
)

type API struct {
	IAPI
	host        string
	port        string
	scheme      string
	accountsURL string
}

var (
	ErrParametersCannotBeNil = fmt.Errorf("parameters cannot be nil")
)

func DefaultAPI() IAPI {
	return &API{host: GetEnv("API_HOST", "localhost"),
		port:        GetEnv("API_PORT", "8080"),
		scheme:      GetEnv("API_SCHEME", "http"),
		accountsURL: "/v1/organisation/accounts"}
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

func buildURL(requestURL string, parameters *url.Values) (*url.URL, error) {
	if parameters == nil {
		return nil, ErrParametersCannotBeNil
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	u.RawQuery = parameters.Encode()

	return u, nil
}

func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
