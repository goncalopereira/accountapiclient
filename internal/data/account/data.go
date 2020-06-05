package account

type LinkData struct {
	TypeOf string `json:"type"`
	ID     string `json:"id"`
}

func newData(id string, organisationID string, attributes Attributes) Account {
	d := Account{OrganisationID: organisationID}
	d.TypeOf = "accounts"
	d.ID = id
	d.Attributes = attributes
	return d
}

type Account struct {
	LinkData
	ResponseOnlyData
	Attributes     `json:"attributes"`
	OrganisationID string `json:"organisation_id"`
	Relationships  `json:"relationships,omitempty"`
}

type ResponseOnlyData struct {
	CreatedOn  string `json:"created_on,omitempty"` //TODO json unmarshal Time later
	ModifiedOn string `json:"modified_on,omitempty"`
	Version    int    `json:"version,omitempty"`
}
