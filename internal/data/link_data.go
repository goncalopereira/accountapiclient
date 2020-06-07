package data

import "github.com/google/uuid"

type LinkData struct {
	TypeOf string    `json:"type"`
	ID     uuid.UUID `json:"id"`
}
