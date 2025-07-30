package utils

import "github.com/google/uuid"

func NewEventID() string {
	return uuid.Must(uuid.NewV7()).String()
}
