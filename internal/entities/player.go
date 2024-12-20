package entities

import "github.com/google/uuid"

type Player struct {
	PublicId uuid.UUID
	Name     string
	Level    int
}
