package repository

import (
	"github.com/google/uuid"
)

func GenerateUuidAsIdTable() (id uuid.UUID) {
	id, _ = uuid.NewV7()
	return
}
