package data

//Data represents the Data response type for Account
// https://api-docs.form3.tech/api.html#organisation-accounts-resource
type Data struct {
	IOutput `json:",omitempty"` //force marshaling to hide fields
	Account `json:"data"`
}
