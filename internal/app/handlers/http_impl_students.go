package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *http) GetStudentsPaginated(ctx *hc.Context) {
	studyPlaceId := ctx.StudyPlaceId()
	query := ctx.PaginationQuery()

	response, err := h.controller.GetStudentsPaginated(ctx, studyPlaceId, &query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) GetStudentById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.GetStudentById(ctx, enrollment, id)
	if err != nil {
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) CreateStudents(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateStudentsRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateStudents(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, response)
}

func (h *http) UpdateStudent(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var request dto.UpdateStudentRequestDTO
	if err = ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.UpdateStudent(ctx, enrollment, id, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteStudentsByIds(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.DeleteStudentsByIdsRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.DeleteStudentsByIds(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
