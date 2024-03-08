package dto

import (
	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
)

type GetTypesByIdRequestDTO models.TypesIds

type CreateGroupEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateGroupsRequestDTO struct {
	List []CreateGroupEntryRequestDTO `json:"list"`
}

type UpdateGroupRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateRoomEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateRoomsRequestDTO struct {
	List []CreateRoomEntryRequestDTO `json:"list"`
}

type UpdateRoomRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateStudentEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateStudentsRequestDTO struct {
	List []CreateStudentEntryRequestDTO `json:"list"`
}

type UpdateStudentRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateSubjectEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateSubjectsRequestDTO struct {
	List []CreateSubjectEntryRequestDTO `json:"list"`
}

type UpdateSubjectRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateTeacherEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateTeachersRequestDTO struct {
	List []CreateTeacherEntryRequestDTO `json:"list"`
}

type UpdateTeacherRequestDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type AddStudentsToGroupRequestDTO struct {
	GroupId    uuid.UUID   `json:"groupId"`
	StudentIds []uuid.UUID `json:"studentIds"`
}

type RemoveStudentFromGroupDTO struct {
	GroupId   uuid.UUID `json:"groupId"`
	StudentId uuid.UUID `json:"studentId"`
}

type AddTutorToGroupsRequestDTO struct {
	GroupIds  []uuid.UUID `json:"groupIds"`
	TeacherId uuid.UUID   `json:"teacherId"`
}

type RemoveGroupTutorRequestDTO struct {
	GroupId   uuid.UUID `json:"groupId"`
	TeacherId uuid.UUID `json:"teacherId"`
}
