package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
)

func CreateRequestData() *account.Data {
	account := NewAccountFromFile("create.json")
	return account
}

func CreateResponseData() *account.Data {
	account := NewAccountFromFile("create-response.json")
	return account
}

func DuplicateAccountErrorResponse() *data.ErrorResponse {
	return &data.ErrorResponse{
		ErrorMessage: "Account cannot be created as it violates a duplicate constraint",
		StatusCode:   http.StatusConflict}
}

func ServerErrorResponse() *data.ErrorResponse {
	errorMessage := NewErrorMessageFromFile("server-error.json")
	errorMessage.StatusCode = 500

	return errorMessage
}
