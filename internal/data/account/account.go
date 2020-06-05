package account

type Account struct {
	LinkData
	ResponseOnlyData
	Attributes     `json:"attributes"`
	OrganisationID string `json:"organisation_id"`
	Relationships  `json:"relationships,omitempty"`
}

func NewAccount(id string, organisationID string, attributes Attributes) Account {
	d := Account{OrganisationID: organisationID}
	d.TypeOf = "accounts"
	d.ID = id
	d.Attributes = attributes

	return d
}
