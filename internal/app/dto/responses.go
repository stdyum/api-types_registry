package dto

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
)

type GroupItemResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Name         string    `json:"name"`
}

type GroupsResponseDTO pagination.Result[GroupItemResponseDTO]

type RoomItemResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Name         string    `json:"name"`
}

type RoomsResponseDTO pagination.Result[RoomItemResponseDTO]

type StudentItemResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Name         string    `json:"name"`
}

type StudentsResponseDTO pagination.Result[StudentItemResponseDTO]

type SubjectItemResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Name         string    `json:"name"`
}

type SubjectsResponseDTO pagination.Result[SubjectItemResponseDTO]

type TeacherItemResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	Name         string    `json:"name"`
}

type TeachersResponseDTO pagination.Result[TeacherItemResponseDTO]
