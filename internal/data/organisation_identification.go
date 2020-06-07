//nolint:lll
package data

type OrganisationIdentification struct {
	Actors                                   []Actor  `json:"actors,omitempty"`
	OrganisationIdentificationIdentification string   `json:"identification,omitempty"`
	OrganisationIdentificationAddress        []string `json:"address,omitempty"` //website example is a string of an array, further work required "[10 Avenue des Champs]"
	OrganisationIdentificationCity           string   `json:"city,omitempty"`
	OrganisationIdentificationCountry        string   `json:"country,omitempty"`
}
