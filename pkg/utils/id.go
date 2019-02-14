package utils

import (
	"log"

	"github.com/google/uuid"
)

// MakeUUID makes UUID string
func MakeUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error while getting id: %v\n", err)
	}
	return id.String()
}
