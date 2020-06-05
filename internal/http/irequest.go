package http

import "fmt"

type IRequest interface {
	Get(reqURL fmt.Stringer) (*Response, error)
	Delete(reqURL fmt.Stringer) (*Response, error)
	Post(reqURL fmt.Stringer, requestData []byte) (*Response, error)
}
