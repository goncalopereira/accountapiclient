package test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"net/http"
)

var (
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
