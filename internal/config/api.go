package config

import (
	"fmt"
	"net/url"
	"os"
)

type API struct {
	IAPI
	apiURL *url.URL
}

var (
	ErrParametersCannotBeNil = fmt.Errorf("parameters cannot be nil")
)

func DefaultAPI() IAPI {
	return &API{apiURL: &url.URL{
		Scheme: GetEnv("API_SCHEME", "http"),
		Host: fmt.Sprintf("%s:%s",
			GetEnv("API_HOST", "localhost"),
			GetEnv("API_PORT", "8080")),
		Path: "/v1/organisation/accounts"}}
}

func (c *API) Accounts(parameters *url.Values) (*url.URL, error) {
	return buildURL(*c.apiURL, parameters)
}

func (c *API) Account(id string, parameters *url.Values) (*url.URL, error) {
	newUrl := *c.apiURL
	newUrl.Path = fmt.Sprintf("%s/%s", c.apiURL.Path, id)
	return buildURL(newUrl, parameters)
}

//makes copy of apiURL and adds parameters to call api.
func buildURL(apiUrl url.URL, parameters *url.Values) (*url.URL, error) {
	if parameters == nil {
		return nil, ErrParametersCannotBeNil
	}

	apiUrl.RawQuery = parameters.Encode()

	return &apiUrl, nil
}

//gets api configuration env variables or default values
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
