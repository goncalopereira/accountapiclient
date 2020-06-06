package account

type Account struct {
	LinkData
	ResponseOnlyData
	Attributes     `json:"attributes"`
	OrganisationID string `json:"organisation_id"`
	Relationships  `json:"relationships,omitempty"`
}
