package data

import "github.com/google/uuid"

type Account struct {
	LinkData
	ResponseOnlyData
	Attributes     `json:"attributes"`
	OrganisationID uuid.UUID `json:"organisation_id"`
	Relationships  `json:"relationships,omitempty"`
}
