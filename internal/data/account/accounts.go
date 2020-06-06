package account

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

type AccountsData struct {
	data.IOutput `json:",omitempty"`
	fmt.Stringer `json:",omitempty"`
	Accounts     *[]Account `json:"data"`
}

func (a AccountsData) String() string {
	account, _ := json.DataToBytes(a)
	return string(account)
}
