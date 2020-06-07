package data

type Relationships struct {
	RelationshipsAccountEvents `json:"account_events,omitempty"`
	RelationshipsMasterAccount `json:"master_account,omitempty"`
}

type RelationshipsAccountEvents struct {
	Data []LinkData `json:"data,omitempty"`
}

type RelationshipsMasterAccount struct {
	Data []LinkData `json:"data,omitempty"`
}
