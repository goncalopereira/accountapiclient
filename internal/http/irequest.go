package http

type IRequest interface {
	Get(reqURL string) (*Response, error)
	Delete(reqURL string) (*Response, error)
	Post(reqURL string, requestData []byte) (*Response, error)
}
