package data

import "github.com/google/uuid"

type Account struct {
	LinkData
	ResponseOnlyData
	Attributes     `json:"attributes"`
	OrganisationID uuid.UUID `json:"organisation_id"`
	Relationships  `json:"relationships,omitempty"`
}

func NewAccount(id uuid.UUID, country string) *Account {
	a := &Account{}
	a.TypeOf = "accounts"
	a.Country = country
	a.ID = id

	return a
}
