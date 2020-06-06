package api

import (
	"fmt"
	"net/url"
	"os"
)

//API holds configuration for client
//env for URL and path
//logic for build account paths.
type API struct {
	apiURL *url.URL
}

var (
	ErrParametersCannotBeNil = fmt.Errorf("parameters cannot be nil")
)

//default values for API based on ENV.
func DefaultAPI() *API {
	return &API{apiURL: &url.URL{
		Scheme: GetEnv("API_SCHEME", "http"),
		Host: fmt.Sprintf("%s:%s",
			GetEnv("API_HOST", "localhost"),
			GetEnv("API_PORT", "8080")),
		Path: "/v1/organisation/accounts"}}
}

//returns full url for accounts resource.
func (c *API) Accounts(parameters *url.Values) (*url.URL, error) {
	return buildURL(*c.apiURL, parameters)
}

//returns full url for given account.
func (c *API) Account(id string, parameters *url.Values) (*url.URL, error) {
	newURL := *c.apiURL
	newURL.Path = fmt.Sprintf("%s/%s", c.apiURL.Path, id)

	return buildURL(newURL, parameters)
}

//makes copy of apiURL and adds parameters to call api.
func buildURL(apiURL url.URL, parameters *url.Values) (*url.URL, error) {
	if parameters == nil {
		return nil, ErrParametersCannotBeNil
	}

	apiURL.RawQuery = parameters.Encode()

	return &apiURL, nil
}

//gets api configuration env variables or default values.
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
