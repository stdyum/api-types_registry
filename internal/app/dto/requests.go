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
	Name string `json:"name"`
}

type DeleteGroupsByIdsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
}

type CreateRoomEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateRoomsRequestDTO struct {
	List []CreateRoomEntryRequestDTO `json:"list"`
}

type UpdateRoomRequestDTO struct {
	Name string `json:"name"`
}

type DeleteRoomsByIdsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
}

type CreateStudentEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateStudentsRequestDTO struct {
	List []CreateStudentEntryRequestDTO `json:"list"`
}

type UpdateStudentRequestDTO struct {
	Name string `json:"name"`
}

type DeleteStudentsByIdsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
}

type CreateSubjectEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateSubjectsRequestDTO struct {
	List []CreateSubjectEntryRequestDTO `json:"list"`
}

type UpdateSubjectRequestDTO struct {
	Name string `json:"name"`
}

type DeleteSubjectsByIdsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
}

type CreateTeacherEntryRequestDTO struct {
	Name string `json:"name"`
}

type CreateTeachersRequestDTO struct {
	List []CreateTeacherEntryRequestDTO `json:"list"`
}

type UpdateTeacherRequestDTO struct {
	Name string `json:"name"`
}

type DeleteTeachersByIdsRequestDTO struct {
	IDs []uuid.UUID `json:"ids"`
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
