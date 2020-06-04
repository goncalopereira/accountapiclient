package config

import "github.com/goncalopereira/accountapiclient/internal/env"

type API struct {
	host   string
	port   string
	scheme string
}

func NewAPI(host string, port string, scheme string) *API {
	return &API{host: host, port: port, scheme: scheme}
}

func DefaultAPI() *API {
	return NewAPI(env.GetEnv("API_HOST", "localhost"), env.GetEnv("API_PORT", "8080"), env.GetEnv("API_SCHEME", "http"))
}

func (c *API) BaseURL() string {
	return c.scheme + "://" + c.host + ":" + c.port
}
