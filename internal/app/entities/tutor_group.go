package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type TutorGroup struct {
	entities.Timed

	StudyPlaceId uuid.UUID
	TeacherId    uuid.UUID
	GroupId      uuid.UUID
}
