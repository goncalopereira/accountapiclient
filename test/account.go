package test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
)

func NewAccountsFromFile(filename string) *account.AccountsData {
	accounts := account.AccountsData{}

	err := json.Unmarshal(ReadJSON(filename), &accounts)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &accounts
}
func NewAccountFromFile(filename string) *account.Data {
	accountToHydrate := account.Data{}

	err := json.Unmarshal(ReadJSON(filename), &accountToHydrate)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &accountToHydrate
}

func NewErrorMessageFromFile(filename string) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}

	err := json.Unmarshal(ReadJSON(filename), &errorResponseToHydrate)
	if err != nil {
		panic("could not hydrate test data")
	}

	return &errorResponseToHydrate
}
