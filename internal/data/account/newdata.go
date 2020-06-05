package account

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Data struct {
	data.IOutput `json:",omitempty"` //force marshaling to hide fields
	fmt.Stringer `json:",omitempty"`
	Account      `json:"data"`
}

func NewAccount(id string, organisationID string, attributes Attributes) *Data {
	a := &Data{}
	a.Account = newData(id, organisationID, attributes)
	return a
}

func NewEmptyAccount() Data {
	return Data{}
}

func (a *Data) String() string {
	account, _ := json.DataToBody(a)
	return string(account)
}
