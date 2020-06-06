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

func (a Data) String() string {
	account, _ := json.DataToBytes(a)
	return string(account)
}
