package handlers

import (
	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-common/http/middlewares"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"google.golang.org/grpc"
)

func (h *http) ConfigureRoutes() *hc.Engine {
	engine := hc.New()
	engine.Use(hc.Recovery())

	v1 := engine.Group("api/v1", hc.Logger(), middlewares.ErrorMiddleware())

	withStudyPlaceId := v1.Group("", middlewares.StudyPlaceMiddleware())

	{
		groupsGroup := withStudyPlaceId.Group("groups")

		groupsGroup.Use(middlewares.AuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetGroupsPaginated)
		groupsGroup.Use(middlewares.AuthMiddleware()).GET("id", h.GetGroupById)
		groupsGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.CreateGroups)
		groupsGroup.Use(middlewares.EnrollmentAuthMiddleware()).PUT(":id", h.UpdateGroup)
		groupsGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.DeleteGroupsByIds)
	}

	{
		roomsGroup := withStudyPlaceId.Group("rooms")

		roomsGroup.Use(middlewares.AuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetRoomsPaginated)
		roomsGroup.Use(middlewares.AuthMiddleware()).GET("id", h.GetRoomById)
		roomsGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.CreateRooms)
		roomsGroup.Use(middlewares.EnrollmentAuthMiddleware()).PUT(":id", h.UpdateRoom)
		roomsGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.DeleteRoomsByIds)
	}

	{
		studentsGroup := withStudyPlaceId.Group("students")

		studentsGroup.Use(middlewares.AuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetStudentsPaginated)
		studentsGroup.Use(middlewares.AuthMiddleware()).GET("id", h.GetStudentById)
		studentsGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.CreateStudents)
		studentsGroup.Use(middlewares.EnrollmentAuthMiddleware()).PUT(":id", h.UpdateStudent)
		studentsGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.DeleteStudentsByIds)
	}

	{
		subjectsGroup := withStudyPlaceId.Group("subjects")

		subjectsGroup.Use(middlewares.AuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetSubjectsPaginated)
		subjectsGroup.Use(middlewares.AuthMiddleware()).GET("id", h.GetSubjectById)
		subjectsGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.CreateSubjects)
		subjectsGroup.Use(middlewares.EnrollmentAuthMiddleware()).PUT(":id", h.UpdateSubject)
		subjectsGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.DeleteSubjectsByIds)
	}

	{
		teachersGroup := withStudyPlaceId.Group("teachers")

		teachersGroup.Use(middlewares.AuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetTeachersPaginated)
		teachersGroup.Use(middlewares.AuthMiddleware()).GET("id", h.GetTeacherById)
		teachersGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.CreateTeachers)
		teachersGroup.Use(middlewares.EnrollmentAuthMiddleware()).PUT(":id", h.UpdateTeacher)
		teachersGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.DeleteTeachersByIds)
	}

	{
		studentsGroupsGroup := withStudyPlaceId.Group("students_groups")

		studentsGroupsGroup.Use(middlewares.EnrollmentAuthMiddleware(), middlewares.PaginationMiddleware(10)).GET("", h.GetStudentsInGroups)
		studentsGroupsGroup.Use(middlewares.EnrollmentAuthMiddleware()).POST("", h.AddStudentsToGroup)
		studentsGroupsGroup.Use(middlewares.EnrollmentAuthMiddleware()).DELETE("", h.RemoveStudentFromGroup)
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	grpcServer := grpc.NewServer()
	types_registry.RegisterTypesRegistryServer(grpcServer, h)
	return grpcServer
}
