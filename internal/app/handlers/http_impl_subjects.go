package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *http) GetSubjectsPaginated(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	query := ctx.PaginationQuery()

	response, err := h.controller.GetSubjectsPaginated(ctx, enrollment, &query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) GetSubjectById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.GetSubjectById(ctx, enrollment, id)
	if err != nil {
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) CreateSubjects(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateSubjectsRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateSubjects(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, response)
}

func (h *http) UpdateSubject(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.UpdateSubjectRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateSubject(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteSubjectById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.DeleteSubjectById(ctx, enrollment, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
