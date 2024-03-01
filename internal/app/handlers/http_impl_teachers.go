package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *http) GetTeachersPaginated(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	query := ctx.PaginationQuery()

	response, err := h.controller.GetTeachersPaginated(ctx, enrollment, &query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) GetTeacherById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.GetTeacherById(ctx, enrollment, id)
	if err != nil {
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) CreateTeachers(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateTeachersRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateTeachers(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, response)
}

func (h *http) UpdateTeacher(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.UpdateTeacherRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateTeacher(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteTeacherById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.DeleteTeacherById(ctx, enrollment, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
