package http

import "net/http"

type IRequest interface {
	Get(req *http.Request) (*Response, error)
	Post(reqURL string, requestData []byte) (*Response, error)
}
