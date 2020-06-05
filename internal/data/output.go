package data

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
)

type Output struct {
	fmt.Stringer
	Account       *account.Account
	ErrorResponse *ErrorResponse
	Accounts      *account.Accounts
}

func (o *Output) String() string {
	if o.ErrorResponse != nil {
		return o.ErrorResponse.String()
	}
	if o.Account != nil {
		return o.Account.String()
	}
	return ""
}
