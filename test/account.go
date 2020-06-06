package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

func NewAccountsFromFile(filename string) *account.AccountsData {
	accounts := account.AccountsData{}
	_ = json.BytesToData(ReadJSON(filename), &accounts)

	return &accounts
}
func NewAccountFromFile(filename string) *account.Data {
	accountToHydrate := account.Data{}
	_ = json.BytesToData(ReadJSON(filename), &accountToHydrate)

	return &accountToHydrate
}

func NewErrorMessageFromFile(filename string) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}
	_ = json.BytesToData(ReadJSON(filename), &errorResponseToHydrate)

	return &errorResponseToHydrate
}
