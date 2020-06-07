package data

//Actor for OrganisationIdentification.
type Actor struct {
	Names     []string `json:"name,omitempty"`
	BirthDate string   `json:"birth_date,omitempty"`
	Residency string   `json:"residency,omitempty"`
}
