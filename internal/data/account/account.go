package account

// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Account struct {
	Data `json:"data"`
}

func NewAccount(id string, organisationID string, attributes Attributes) *Account {
	a := &Account{}
	a.Data = newData(id, organisationID, attributes)
	return a
}

func NewEmptyAccount() Account {
	return Account{}
}
