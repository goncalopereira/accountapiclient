package data

type ErrorResponse struct {
	IOutput      `json:",omitempty"`
	ErrorMessage string `json:"error_message"`
	StatusCode   int
}
