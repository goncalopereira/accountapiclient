package test

import (
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

func NewAccountFromFile(filename string) *account.Account {
	accountToHydrate := account.NewEmptyAccount()
	_ = json.BodyToData(ReadJSON(filename), &accountToHydrate)
	return &accountToHydrate
}
