package http

import "net/http"

type IRequest interface {
	Do(req *http.Request) (*Response, error)
}
