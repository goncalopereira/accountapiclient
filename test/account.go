package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

func NewAccountsFromFile(filename string) *account.AccountsData {
	accounts := account.AccountsData{}
	err := json.BytesToData(ReadJSON(filename), &accounts)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &accounts
}
func NewAccountFromFile(filename string) *account.Data {
	accountToHydrate := account.Data{}
	err := json.BytesToData(ReadJSON(filename), &accountToHydrate)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &accountToHydrate
}

func NewErrorMessageFromFile(filename string) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}
	err := json.BytesToData(ReadJSON(filename), &errorResponseToHydrate)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &errorResponseToHydrate
}
