//nolint:lll
package account

type PrivateIdentification struct {
	PrivateIdentificationBirthDate      string   `json:"birth_date,omitempty"`
	PrivateIdentificationBirthCountry   string   `json:"birth_country,omitempty"`
	PrivateIdentificationIdentification string   `json:"identification,omitempty"`
	PrivateIdentificationAddress        []string `json:"address,omitempty"` //typo in Fetch example shows up as string   "address": "[10 Avenue des Champs]",
	PrivateIdentificationCity           string   `json:"city,omitempty"`
	PrivateIdentificationCountry        string   `json:"country,omitempty"`
}
