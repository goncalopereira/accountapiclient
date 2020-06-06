package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"net/http"
)

func AccountCreateRequest() *account.Data {
	account := NewAccountFromFile("create.json")
	return account
}

func AccountCreateResponse() *account.Data {
	account := NewAccountFromFile("create-response.json")
	return account
}

func DuplicateAccountErrorResponse() *data.ErrorResponse {
	return &data.ErrorResponse{
		ErrorMessage: "Account cannot be created as it violates a duplicate constraint",
		StatusCode:   http.StatusConflict}
}
