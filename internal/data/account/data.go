package account

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
)

// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Data struct {
	data.IOutput `json:",omitempty"` //force marshaling to hide fields
	Account      `json:"data"`
}
