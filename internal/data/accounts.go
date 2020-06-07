package data

//AccountsData represents the Data level type from Response with an Account array.
type AccountsData struct {
	IOutput  `json:",omitempty"`
	Accounts *[]Account `json:"data"`
}
