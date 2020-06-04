package http

type IRequest interface {
	Get(endpoint string) (*Response, error)
	Post(endpoint string, requestData []byte) (*Response, error)
}
