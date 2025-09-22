package domain

import "github.com/google/uuid"

type Player struct {
	PublicID uuid.UUID `json:"publicId"`
	Name     string    `json:"name"`
	Level    int       `json:"level"`
}
