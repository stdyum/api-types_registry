package handlers

import (
	"github.com/stdyum/api-common/hc"
	confHttp "github.com/stdyum/api-common/http"
	"github.com/stdyum/api-types-registry/internal/app/controllers"
)

type HTTP interface {
	confHttp.Routes

	GetGroupsPaginated(ctx *hc.Context)
	GetGroupById(ctx *hc.Context)
	CreateGroups(ctx *hc.Context)
	UpdateGroup(ctx *hc.Context)
	DeleteGroupsByIds(ctx *hc.Context)

	GetRoomsPaginated(ctx *hc.Context)
	GetRoomById(ctx *hc.Context)
	CreateRooms(ctx *hc.Context)
	UpdateRoom(ctx *hc.Context)
	DeleteRoomsByIds(ctx *hc.Context)

	GetSubjectsPaginated(ctx *hc.Context)
	GetSubjectById(ctx *hc.Context)
	CreateSubjects(ctx *hc.Context)
	UpdateSubject(ctx *hc.Context)
	DeleteSubjectsByIds(ctx *hc.Context)

	GetTeachersPaginated(ctx *hc.Context)
	GetTeacherById(ctx *hc.Context)
	CreateTeachers(ctx *hc.Context)
	UpdateTeacher(ctx *hc.Context)
	DeleteTeachersByIds(ctx *hc.Context)

	GetStudentsPaginated(ctx *hc.Context)
	GetStudentById(ctx *hc.Context)
	CreateStudents(ctx *hc.Context)
	UpdateStudent(ctx *hc.Context)
	DeleteStudentsByIds(ctx *hc.Context)

	GetStudentsInGroups(ctx *hc.Context)
	AddStudentsToGroup(ctx *hc.Context)
	RemoveStudentFromGroup(ctx *hc.Context)

	GetTeacherTuitionGroups(ctx *hc.Context)
	AddTutorToGroups(ctx *hc.Context)
	RemoveGroupTutor(ctx *hc.Context)
}

type http struct {
	controller controllers.Controller
}

func NewHTTP(controller controllers.Controller) HTTP {
	return &http{
		controller: controller,
	}
}
