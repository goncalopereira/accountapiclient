package account

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

type Accounts struct {
	data.IOutput `json:",omitempty"`
	fmt.Stringer `json:",omitempty"`
	Data         *[]Data `json:"data"`
}

func (a *Accounts) String() string {
	account, _ := json.DataToBody(a)
	return string(account)
}
