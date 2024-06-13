package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *http) GetStudentsInGroups(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	query := ctx.PaginationQuery()

	response, err := h.controller.GetStudentsInGroupsPaginated(ctx, enrollment, &query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) AddStudentsToGroup(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.AddStudentsToGroupRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.AddStudentsToGroup(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusCreated)
}

func (h *http) RemoveStudentFromGroup(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.RemoveStudentFromGroupDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	err := h.controller.RemoveStudentFromGroup(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusCreated)
}

func (h *http) GetTeacherTuitionGroups(ctx *hc.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *http) AddTutorToGroups(ctx *hc.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *http) RemoveGroupTutor(ctx *hc.Context) {
	//TODO implement me
	panic("implement me")
}
