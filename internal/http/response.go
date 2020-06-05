package http

type Response struct {
	StatusCode int
	Body       []byte
}

func NewResponse(statusCode int, body []byte) *Response {
	return &Response{StatusCode: statusCode, Body: body}
}
