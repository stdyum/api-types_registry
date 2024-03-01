package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"github.com/stdyum/api-types-registry/internal/app/dto"
	"github.com/stdyum/api-types-registry/internal/app/entities"
	"github.com/stdyum/api-types-registry/internal/app/repositories"
)

type Controller interface {
	GetGroupsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.GroupsResponseDTO, error)
	GetGroupsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.GroupItemResponseDTO, error)
	GetGroupById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.GroupItemResponseDTO, error)
	CreateGroups(ctx context.Context, enrollment models.Enrollment, request dto.CreateGroupsRequestDTO) ([]entities.Group, error)
	UpdateGroup(ctx context.Context, enrollment models.Enrollment, request dto.UpdateGroupRequestDTO) error
	DeleteGroupById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	GetRoomsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.RoomsResponseDTO, error)
	GetRoomsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.RoomItemResponseDTO, error)
	GetRoomById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.RoomItemResponseDTO, error)
	CreateRooms(ctx context.Context, enrollment models.Enrollment, request dto.CreateRoomsRequestDTO) ([]entities.Room, error)
	UpdateRoom(ctx context.Context, enrollment models.Enrollment, request dto.UpdateRoomRequestDTO) error
	DeleteRoomById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	GetSubjectsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.SubjectsResponseDTO, error)
	GetSubjectsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.SubjectItemResponseDTO, error)
	GetSubjectById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.SubjectItemResponseDTO, error)
	CreateSubjects(ctx context.Context, enrollment models.Enrollment, request dto.CreateSubjectsRequestDTO) ([]entities.Subject, error)
	UpdateSubject(ctx context.Context, enrollment models.Enrollment, request dto.UpdateSubjectRequestDTO) error
	DeleteSubjectById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	GetTeachersPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.TeachersResponseDTO, error)
	GetTeachersByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.TeacherItemResponseDTO, error)
	GetTeacherById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.TeacherItemResponseDTO, error)
	CreateTeachers(ctx context.Context, enrollment models.Enrollment, request dto.CreateTeachersRequestDTO) ([]entities.Teacher, error)
	UpdateTeacher(ctx context.Context, enrollment models.Enrollment, request dto.UpdateTeacherRequestDTO) error
	DeleteTeacherById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	GetStudentsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.StudentsResponseDTO, error)
	GetStudentsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.StudentItemResponseDTO, error)
	GetStudentById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.StudentItemResponseDTO, error)
	CreateStudents(ctx context.Context, enrollment models.Enrollment, request dto.CreateStudentsRequestDTO) ([]entities.Student, error)
	UpdateStudent(ctx context.Context, enrollment models.Enrollment, request dto.UpdateStudentRequestDTO) error
	DeleteStudentById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error

	GetStudentsInGroup(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.StudentItemResponseDTO, error)
	AddStudentsToGroup(ctx context.Context, enrollment models.Enrollment, request dto.AddStudentsToGroupRequestDTO) error
	RemoveStudentFromGroup(ctx context.Context, enrollment models.Enrollment, request dto.RemoveStudentFromGroupDTO) error

	GetTeacherTuitionGroups(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.GroupItemResponseDTO, error)
	AddTutorToGroups(ctx context.Context, enrollment models.Enrollment, request dto.AddTutorToGroupsRequestDTO) error
	RemoveGroupTutor(ctx context.Context, enrollment models.Enrollment, request dto.RemoveGroupTutorRequestDTO) error
}

type controller struct {
	repository repositories.Repository
	client     studyplaces.StudyplacesClient
}

func New(repository repositories.Repository, client studyplaces.StudyplacesClient) Controller {
	return &controller{
		repository: repository,
		client:     client,
	}
}
