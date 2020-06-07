package data

type AccountsData struct {
	IOutput  `json:",omitempty"`
	Accounts *[]Account `json:"data"`
}
