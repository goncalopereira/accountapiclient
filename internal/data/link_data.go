package data

import "github.com/google/uuid"

//LinkData represents Type and ID for multiple types.
type LinkData struct {
	TypeOf string    `json:"type"`
	ID     uuid.UUID `json:"id"`
}
