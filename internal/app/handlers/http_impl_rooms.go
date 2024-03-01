package handlers

import (
	netHttp "net/http"

	"github.com/stdyum/api-common/hc"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *http) GetRoomsPaginated(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	query := ctx.PaginationQuery()

	response, err := h.controller.GetRoomsPaginated(ctx, enrollment, &query)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) GetRoomById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.GetRoomById(ctx, enrollment, id)
	if err != nil {
		return
	}

	ctx.JSON(netHttp.StatusOK, response)
}

func (h *http) CreateRooms(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.CreateRoomsRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	response, err := h.controller.CreateRooms(ctx, enrollment, request)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(netHttp.StatusCreated, response)
}

func (h *http) UpdateRoom(ctx *hc.Context) {
	enrollment := ctx.Enrollment()

	var request dto.UpdateRoomRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := h.controller.UpdateRoom(ctx, enrollment, request); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}

func (h *http) DeleteRoomById(ctx *hc.Context) {
	enrollment := ctx.Enrollment()
	id, err := ctx.UUIDParam("id")
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if err = h.controller.DeleteRoomById(ctx, enrollment, id); err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.Status(netHttp.StatusNoContent)
}
