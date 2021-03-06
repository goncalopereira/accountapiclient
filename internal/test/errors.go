package test

import (
	"fmt"
	"net/http"

	"github.com/goncalopereira/accountapiclient/internal/data"
)

var (
	//ErrBrokenHTTPClient represents error message from test issues with Request
	ErrBrokenHTTPClient = fmt.Errorf("broken http connection")
)

//ServerErrorResponse returns valid static ErrorResponse for generic 500 Server Error.
func ServerErrorResponse() *data.ErrorResponse {
	errorMessage := NewErrorMessageFromFile("server-error.json", http.StatusInternalServerError)

	return errorMessage
}

//DuplicateAccountErrorResponse returns valid static ErrorResponse
//for existing Account error.
func DuplicateAccountErrorResponse() *data.ErrorResponse {
	return &data.ErrorResponse{
		ErrorMessage: "Account cannot be created as it violates a duplicate constraint",
		StatusCode:   http.StatusConflict}
}
