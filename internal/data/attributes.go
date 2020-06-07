package data

type Attributes struct {
	PrivateIdentification      `json:"private_identification,omitempty"`
	OrganisationIdentification `json:"organisation_identification,omitempty"`
	Country                    string   `json:"country"`
	BaseCurrency               string   `json:"base_currency,omitempty"`
	BankID                     string   `json:"bank_id,omitempty"`
	BankIDCode                 string   `json:"bank_id_code,omitempty"`
	AccountNumber              string   `json:"account_number,omitempty"`
	Bic                        string   `json:"bic,omitempty"`
	Iban                       string   `json:"iban,omitempty"`
	CustomerID                 string   `json:"customer_id,omitempty"`
	Name                       []string `json:"name,omitempty"`
	AlternativeNames           []string `json:"alternative_names,omitempty"`
	AccountClassification      string   `json:"account_classification,omitempty"`
	JointAccount               bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut      bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification    string   `json:"secondary_identification,omitempty"`
	Switched                   bool     `json:"switched,omitempty"`
	Status                     string   `json:"status,omitempty"`
}
