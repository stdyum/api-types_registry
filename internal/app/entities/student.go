package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type Student struct {
	entities.Timed

	ID           uuid.UUID
	StudyPlaceId uuid.UUID
	Name         string
}
