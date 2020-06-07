package data

//ErrorResponse is one of the valid types of Output
//represents known behaviors that are not the valid result of a command
//it does not use the error interface to separate flow logic from underlying errors like http connection.
type ErrorResponse struct {
	IOutput      `json:",omitempty"`
	ErrorMessage string `json:"error_message"`
	StatusCode   int
}
