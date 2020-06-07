package data

//Relationships for Account.
type Relationships struct {
	RelationshipsAccountEvents `json:"account_events,omitempty"`
	RelationshipsMasterAccount `json:"master_account,omitempty"`
}

//RelationshipsAccountEvents AccountEvents for Account.
type RelationshipsAccountEvents struct {
	Data []LinkData `json:"data,omitempty"`
}

//RelationshipsMasterAccount MasterAccount for Account.
type RelationshipsMasterAccount struct {
	Data []LinkData `json:"data,omitempty"`
}
