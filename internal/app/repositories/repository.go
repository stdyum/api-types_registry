package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

type Repository interface {
	GetGroupsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Group, int, error)
	GetGroupsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Group, error)
	GetGroupById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Group, error)
	CreateGroups(ctx context.Context, groups []entities.Group) error
	UpdateGroup(ctx context.Context, group entities.Group) error
	DeleteGroupsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error

	GetRoomsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Room, int, error)
	GetRoomsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Room, error)
	GetRoomById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Room, error)
	CreateRooms(ctx context.Context, rooms []entities.Room) error
	UpdateRoom(ctx context.Context, room entities.Room) error
	DeleteRoomsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error

	GetSubjectsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Subject, int, error)
	GetSubjectsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Subject, error)
	GetSubjectById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Subject, error)
	CreateSubjects(ctx context.Context, subjects []entities.Subject) error
	UpdateSubject(ctx context.Context, subject entities.Subject) error
	DeleteSubjectsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error

	GetTeachersPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Teacher, int, error)
	GetTeachersByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Teacher, error)
	GetTeacherById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Teacher, error)
	CreateTeachers(ctx context.Context, teachers []entities.Teacher) error
	UpdateTeacher(ctx context.Context, teacher entities.Teacher) error
	DeleteTeachersByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error

	GetStudentsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Student, int, error)
	GetStudentsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Student, error)
	GetStudentById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Student, error)
	CreateStudents(ctx context.Context, students []entities.Student) error
	UpdateStudent(ctx context.Context, student entities.Student) error
	DeleteStudentsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error

	GetStudentsInGroup(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID) ([]entities.Student, error)
	AddStudentsToGroup(ctx context.Context, studyPlaceId uuid.UUID, studentIds []uuid.UUID, groupId uuid.UUID) error
	RemoveStudentFromGroup(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID, studentId uuid.UUID) error

	GetTeacherTuitionGroups(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID) ([]entities.Group, error)
	AddTutorToGroups(ctx context.Context, studyPlaceId uuid.UUID, groupIds []uuid.UUID, teacherId uuid.UUID) error
	RemoveGroupTutor(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID, teacherId uuid.UUID) error
}

type repository struct {
	database *sql.DB
}

func New(database *sql.DB) Repository {
	return &repository{
		database: database,
	}
}
