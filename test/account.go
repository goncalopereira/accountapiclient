package test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
)

const ErrCouldNotHydrateTestData = "could not hydrate test data"

func NewAccountsFromFile(filename string) *account.AccountsData {
	accounts := account.AccountsData{}

	err := json.Unmarshal(ReadJSON(filename), &accounts)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	return &accounts
}
func NewAccountFromFile(filename string) *account.Data {
	accountToHydrate := account.Data{}

	err := json.Unmarshal(ReadJSON(filename), &accountToHydrate)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	return &accountToHydrate
}

func NewErrorMessageFromFile(filename string) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}

	err := json.Unmarshal(ReadJSON(filename), &errorResponseToHydrate)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	return &errorResponseToHydrate
}
