package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-types-registry/internal/app/dto"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (c *controller) GetRoomsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.RoomsResponseDTO, error) {
	rooms, amount, err := c.repository.GetRoomsPaginated(ctx, enrollment.StudyPlaceId, paginationQuery)
	if err != nil {
		return dto.RoomsResponseDTO{}, err
	}

	paginationResult := pagination.FromArrayAndAmount(rooms, amount, paginationQuery,
		func(el entities.Room) dto.RoomItemResponseDTO {
			return dto.RoomItemResponseDTO{
				ID:           el.ID,
				StudyPlaceId: el.StudyPlaceId,
				Name:         el.Name,
			}
		},
	)

	return dto.RoomsResponseDTO(paginationResult), nil
}

func (c *controller) GetRoomsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.RoomItemResponseDTO, error) {
	rooms, err := c.repository.GetRoomsByIds(ctx, enrollment.StudyPlaceId, ids)
	if err != nil {
		return nil, err
	}

	response := uslices.MapFunc(rooms, func(item entities.Room) dto.RoomItemResponseDTO {
		return dto.RoomItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	})

	return response, nil
}

func (c *controller) GetRoomById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.RoomItemResponseDTO, error) {
	room, err := c.repository.GetRoomById(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.RoomItemResponseDTO{}, err
	}

	return dto.RoomItemResponseDTO{
		ID:           room.ID,
		StudyPlaceId: room.StudyPlaceId,
		Name:         room.Name,
	}, nil
}

func (c *controller) CreateRooms(ctx context.Context, enrollment models.Enrollment, request dto.CreateRoomsRequestDTO) ([]entities.Room, error) {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return nil, err
	}

	rooms := uslices.MapFunc(request.List, func(item dto.CreateRoomEntryRequestDTO) entities.Room {
		return entities.Room{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Name:         item.Name,
		}
	})

	if err := c.repository.CreateRooms(ctx, rooms); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (c *controller) UpdateRoom(ctx context.Context, enrollment models.Enrollment, request dto.UpdateRoomRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	room := entities.Room{
		ID:           request.ID,
		StudyPlaceId: enrollment.StudyPlaceId,
		Name:         request.Name,
	}

	return c.repository.UpdateRoom(ctx, room)
}

func (c *controller) DeleteRoomById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.DeleteRoomById(ctx, enrollment.StudyPlaceId, id)
}
