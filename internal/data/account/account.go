package account

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Account struct {
	data.IOutput `json:",omitempty"` //force marshaling to hide fields
	fmt.Stringer `json:",omitempty"`
	Data         `json:"data"`
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
