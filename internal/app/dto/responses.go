package dto

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
)

type TypesModelsResponseDTO struct {
	GroupsIds   map[string]TypesModelsGroupResponseDTO   `json:"groupsIds"`
	RoomsIds    map[string]TypesModelsRoomResponseDTO    `json:"roomsIds"`
	StudentsIds map[string]TypesModelsStudentResponseDTO `json:"studentsIds"`
	SubjectsIds map[string]TypesModelsSubjectResponseDTO `json:"subjectsIds"`
	TeachersIds map[string]TypesModelsTeacherResponseDTO `json:"teachersIds"`
}

type TypesModelsGroupResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TypesModelsRoomResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TypesModelsStudentResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TypesModelsSubjectResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TypesModelsTeacherResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

type StudentInGroupResponseDTO struct {
	StudyPlaceId uuid.UUID `json:"studyPlaceId"`
	StudentId    uuid.UUID `json:"studentId"`
	StudentName  string    `json:"studentName"`
	GroupId      uuid.UUID `json:"groupId"`
	GroupName    string    `json:"groupName"`
}

type StudentsInGroupsResponseDTO pagination.Result[StudentInGroupResponseDTO]
