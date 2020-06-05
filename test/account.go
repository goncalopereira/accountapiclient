package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

func NewAccountFromFile(filename string) *account.Account {
	accountToHydrate := account.NewEmptyAccount()
	_ = json.BodyToData(ReadJSON(filename), &accountToHydrate)
	return &accountToHydrate
}

func NewErrorMessageFromFile(filename string) *data.ErrorResponse {
	errorResponseToHydrate := data.ErrorResponse{}
	_ = json.BodyToData(ReadJSON(filename), &errorResponseToHydrate)
	return &errorResponseToHydrate
}
