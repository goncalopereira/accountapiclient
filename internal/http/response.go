package http

//Response wraps the http.Response StatusCode and copied Body
//hides http.Response handling during mock.
type Response struct {
	StatusCode int
	Body       []byte
}
