package data

//ResponseOnlyData has fields that only appear during Response.
type ResponseOnlyData struct {
	CreatedOn  string `json:"created_on,omitempty"` //json unmarshal Time some other day
	ModifiedOn string `json:"modified_on,omitempty"`
	Version    int    `json:"version,omitempty"`
}
