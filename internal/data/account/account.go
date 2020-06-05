package account

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Account struct {
	fmt.Stringer
	Data `json:"data"`
}

type Accounts struct {
	fmt.Stringer
	Data *[]Data `json:"data"`
}

func (a *Accounts) String() string {
	account, _ := json.DataToBody(a)
	return string(account)
}

func NewAccount(id string, organisationID string, attributes Attributes) *Account {
	a := &Account{}
	a.Data = newData(id, organisationID, attributes)
	return a
}

func NewEmptyAccount() Account {
	return Account{}
}

func (a *Account) String() string {
	account, _ := json.DataToBody(a)
	return string(account)
}
