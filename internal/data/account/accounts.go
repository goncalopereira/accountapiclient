package account

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
)

type AccountsData struct {
	data.IOutput `json:",omitempty"`
	Accounts     *[]Account `json:"data"`
}
