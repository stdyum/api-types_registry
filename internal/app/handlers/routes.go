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

	{
		groupsGroup := v1.Group("groups").Use(middlewares.EnrollmentAuthMiddleware())

		groupsGroup.Use(middlewares.PaginationMiddleware(10)).GET("", h.GetGroupsPaginated)
		groupsGroup.GET("id", h.GetGroupById)
		groupsGroup.POST("", h.CreateGroups)
		groupsGroup.PUT("", h.UpdateGroup)
		groupsGroup.DELETE("", h.DeleteGroupsByIds)
	}

	{
		roomsGroup := v1.Group("rooms").Use(middlewares.EnrollmentAuthMiddleware())

		roomsGroup.Use(middlewares.PaginationMiddleware(10)).GET("", h.GetRoomsPaginated)
		roomsGroup.GET("id", h.GetRoomById)
		roomsGroup.POST("", h.CreateRooms)
		roomsGroup.PUT("", h.UpdateRoom)
		roomsGroup.DELETE("", h.DeleteRoomsByIds)
	}

	{
		studentsGroup := v1.Group("students").Use(middlewares.EnrollmentAuthMiddleware())

		studentsGroup.Use(middlewares.PaginationMiddleware(10)).GET("", h.GetStudentsPaginated)
		studentsGroup.GET("id", h.GetStudentById)
		studentsGroup.POST("", h.CreateStudents)
		studentsGroup.PUT("", h.UpdateStudent)
		studentsGroup.DELETE("", h.DeleteStudentsByIds)
	}

	{
		subjectsGroup := v1.Group("subjects").Use(middlewares.EnrollmentAuthMiddleware())

		subjectsGroup.Use(middlewares.PaginationMiddleware(10)).GET("", h.GetSubjectsPaginated)
		subjectsGroup.GET("id", h.GetSubjectById)
		subjectsGroup.POST("", h.CreateSubjects)
		subjectsGroup.PUT("", h.UpdateSubject)
		subjectsGroup.DELETE("", h.DeleteSubjectsByIds)
	}

	{
		teachersGroup := v1.Group("teachers").Use(middlewares.EnrollmentAuthMiddleware())

		teachersGroup.Use(middlewares.PaginationMiddleware(10)).GET("", h.GetTeachersPaginated)
		teachersGroup.GET("id", h.GetTeacherById)
		teachersGroup.POST("", h.CreateTeachers)
		teachersGroup.PUT("", h.UpdateTeacher)
		teachersGroup.DELETE("", h.DeleteTeachersByIds)
	}

	return engine
}

func (h *gRPC) ConfigureRoutes() *grpc.Server {
	grpcServer := grpc.NewServer()
	types_registry.RegisterTypesRegistryServer(grpcServer, h)
	return grpcServer
}
