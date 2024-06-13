package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/studyplaces"
	"github.com/stdyum/api-types-registry/internal/app/dto"
	"github.com/stdyum/api-types-registry/internal/app/repositories"
)

type Controller interface {
	GetTypesById(ctx context.Context, enrollment models.Enrollment, ids dto.GetTypesByIdRequestDTO) (dto.TypesModelsResponseDTO, error)

	GetGroupsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.GroupsResponseDTO, error)
	GetGroupsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.GroupItemResponseDTO, error)
	GetGroupById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.GroupItemResponseDTO, error)
	CreateGroups(ctx context.Context, enrollment models.Enrollment, request dto.CreateGroupsRequestDTO) ([]dto.GroupItemResponseDTO, error)
	UpdateGroup(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateGroupRequestDTO) error
	DeleteGroupsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteGroupsByIdsRequestDTO) error

	GetRoomsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.RoomsResponseDTO, error)
	GetRoomsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.RoomItemResponseDTO, error)
	GetRoomById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.RoomItemResponseDTO, error)
	CreateRooms(ctx context.Context, enrollment models.Enrollment, request dto.CreateRoomsRequestDTO) ([]dto.RoomItemResponseDTO, error)
	UpdateRoom(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateRoomRequestDTO) error
	DeleteRoomsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteRoomsByIdsRequestDTO) error

	GetSubjectsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.SubjectsResponseDTO, error)
	GetSubjectsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.SubjectItemResponseDTO, error)
	GetSubjectById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.SubjectItemResponseDTO, error)
	CreateSubjects(ctx context.Context, enrollment models.Enrollment, request dto.CreateSubjectsRequestDTO) ([]dto.SubjectItemResponseDTO, error)
	UpdateSubject(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateSubjectRequestDTO) error
	DeleteSubjectsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteSubjectsByIdsRequestDTO) error

	GetTeachersPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.TeachersResponseDTO, error)
	GetTeachersByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.TeacherItemResponseDTO, error)
	GetTeacherById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.TeacherItemResponseDTO, error)
	CreateTeachers(ctx context.Context, enrollment models.Enrollment, request dto.CreateTeachersRequestDTO) ([]dto.TeacherItemResponseDTO, error)
	UpdateTeacher(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateTeacherRequestDTO) error
	DeleteTeachersByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteTeachersByIdsRequestDTO) error

	GetStudentsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.StudentsResponseDTO, error)
	GetStudentsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.StudentItemResponseDTO, error)
	GetStudentById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.StudentItemResponseDTO, error)
	CreateStudents(ctx context.Context, enrollment models.Enrollment, request dto.CreateStudentsRequestDTO) ([]dto.StudentItemResponseDTO, error)
	UpdateStudent(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateStudentRequestDTO) error
	DeleteStudentsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteStudentsByIdsRequestDTO) error

	GetStudentsInGroup(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.StudentItemResponseDTO, error)
	GetStudentsInGroupsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery *pagination.CreatedAtPageQuery) (dto.StudentsInGroupsResponseDTO, error)
	GetStudentGroups(ctx context.Context, enrollment models.Enrollment, studentId uuid.UUID) ([]dto.GroupItemResponseDTO, error)
	AddStudentsToGroup(ctx context.Context, enrollment models.Enrollment, request dto.AddStudentsToGroupRequestDTO) error
	RemoveStudentFromGroup(ctx context.Context, enrollment models.Enrollment, request dto.RemoveStudentFromGroupDTO) error

	GetTeacherTuitionGroups(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.GroupItemResponseDTO, error)
	AddTutorToGroups(ctx context.Context, enrollment models.Enrollment, request dto.AddTutorToGroupsRequestDTO) error
	RemoveGroupTutor(ctx context.Context, enrollment models.Enrollment, request dto.RemoveGroupTutorRequestDTO) error

	GetGroupIdsWithStudents(ctx context.Context, enrollment models.Enrollment) ([]uuid.UUID, error)
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
