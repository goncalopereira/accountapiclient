package test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
)

const ErrCouldNotHydrateTestData = "could not hydrate test data"

//NewAccountsFromFile reads relative file path and returns AccountsData ([]Account).
func NewAccountsFromFile(filename string) *account.AccountsData {
	accounts := account.AccountsData{}

	err := json.Unmarshal(ReadJSON(filename), &accounts)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	return &accounts
}

//NewAccountsFromFile reads relative file path and returns Data (Account).
func NewAccountFromFile(filename string) *account.Data {
	accountToHydrate := account.Data{}

	err := json.Unmarshal(ReadJSON(filename), &accountToHydrate)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	return &accountToHydrate
}

//NewAccountsFromFile reads relative file path and http.StatusCode and returns ErrorResponse.
func NewErrorMessageFromFile(filename string, statusCode int) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}

	err := json.Unmarshal(ReadJSON(filename), &errorResponseToHydrate)
	if err != nil {
		panic(ErrCouldNotHydrateTestData)
	}

	errorResponseToHydrate.StatusCode = statusCode

	return &errorResponseToHydrate
}
