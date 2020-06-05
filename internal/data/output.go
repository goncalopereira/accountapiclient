package data

import "github.com/goncalopereira/accountapiclient/internal/data/account"

type Output struct {
	Account       *account.Account
	ErrorResponse *ErrorResponse
	Accounts      *[]account.Account
}
