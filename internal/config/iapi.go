package config

import "net/url"

type IAPI interface {
	Accounts(parameters *url.Values) (*url.URL, error)
	Account(id string, parameters *url.Values) (*url.URL, error)
}
