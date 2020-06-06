package http

import "net/http"

type IRequest interface {
	Get(req *http.Request) (*Response, error)
}
