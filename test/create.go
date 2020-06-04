package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
)

//func AccountCreateQuery() *command.Query {
//	account := NewAccountFromFile("create.json")
//	query := command.Query{}
//	query.ID = account.ID
//	return &query
//}

func AccountCreateResponse() *account.Account {
	account := NewAccountFromFile("create-response.json")
	return account
}

func DuplicateAccountErrorResponse() *data.ErrorResponse {
	return &data.ErrorResponse{ErrorMessage: "Account cannot be created as it violates a duplicate constraint"}
}
