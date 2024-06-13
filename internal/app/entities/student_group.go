package entities

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/entities"
)

type StudentGroup struct {
	entities.Timed

	StudyPlaceId uuid.UUID
	StudentId    uuid.UUID
	GroupId      uuid.UUID
}

type AggregatedStudentGroup struct {
	entities.Timed

	StudyPlaceId uuid.UUID
	StudentId    uuid.UUID
	Student      string
	GroupId      uuid.UUID
	Group        string
}
