package api

import (
	"fmt"
	"github.com/google/uuid"
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

//DefaultAPI returns an API with default env config.
func DefaultAPI() *API {
	return &API{apiURL: &url.URL{
		Scheme: getEnv("API_SCHEME", "http"),
		Host: fmt.Sprintf("%s:%s",
			getEnv("API_HOST", "localhost"),
			getEnv("API_PORT", "8080")),
		Path: "/v1/organisation/accounts"}}
}

//Accounts returns the full url for Account resources.
func (c *API) Accounts(parameters *url.Values) (*url.URL, error) {
	return buildURL(*c.apiURL, parameters)
}

//Accounts returns the full url for a specific Account based on id.
func (c *API) Account(id uuid.UUID, parameters *url.Values) (*url.URL, error) {
	newURL := *c.apiURL
	newURL.Path = fmt.Sprintf("%s/%s", c.apiURL.Path, id)

	return buildURL(newURL, parameters)
}

func buildURL(apiURL url.URL, parameters *url.Values) (*url.URL, error) {
	if parameters == nil {
		return nil, ErrParametersCannotBeNil
	}

	apiURL.RawQuery = parameters.Encode()

	return &apiURL, nil
}

//gets api configuration env variables or default values.
func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
