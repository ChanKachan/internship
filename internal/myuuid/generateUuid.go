package myuuid

import (
	"github.com/google/uuid"
)

func GenerateUuid() uuid.UUID {
	return uuid.New()
}
