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

func (c *controller) GetGroupsPaginated(ctx context.Context, enrollment models.Enrollment, paginationQuery pagination.Query) (dto.GroupsResponseDTO, error) {
	groups, amount, err := c.repository.GetGroupsPaginated(ctx, enrollment.StudyPlaceId, paginationQuery)
	if err != nil {
		return dto.GroupsResponseDTO{}, err
	}

	paginationResult := pagination.FromArrayAndAmount(groups, amount, paginationQuery,
		func(el entities.Group) dto.GroupItemResponseDTO {
			return dto.GroupItemResponseDTO{
				ID:           el.ID,
				StudyPlaceId: el.StudyPlaceId,
				Name:         el.Name,
			}
		},
	)

	return dto.GroupsResponseDTO(paginationResult), nil
}

func (c *controller) GetGroupsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.GroupItemResponseDTO, error) {
	groups, err := c.repository.GetGroupsByIds(ctx, enrollment.StudyPlaceId, ids)
	if err != nil {
		return nil, err
	}

	response := uslices.MapFunc(groups, func(item entities.Group) dto.GroupItemResponseDTO {
		return dto.GroupItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	})

	return response, nil
}

func (c *controller) GetGroupById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.GroupItemResponseDTO, error) {
	group, err := c.repository.GetGroupById(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.GroupItemResponseDTO{}, err
	}

	return dto.GroupItemResponseDTO{
		ID:           group.ID,
		StudyPlaceId: group.StudyPlaceId,
		Name:         group.Name,
	}, nil
}

func (c *controller) CreateGroups(ctx context.Context, enrollment models.Enrollment, request dto.CreateGroupsRequestDTO) ([]entities.Group, error) {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return nil, err
	}

	groups := uslices.MapFunc(request.List, func(item dto.CreateGroupEntryRequestDTO) entities.Group {
		return entities.Group{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Name:         item.Name,
		}
	})

	if err := c.repository.CreateGroups(ctx, groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func (c *controller) UpdateGroup(ctx context.Context, enrollment models.Enrollment, request dto.UpdateGroupRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	group := entities.Group{
		ID:           request.ID,
		StudyPlaceId: enrollment.StudyPlaceId,
		Name:         request.Name,
	}

	return c.repository.UpdateGroup(ctx, group)
}

func (c *controller) DeleteGroupById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.DeleteGroupById(ctx, enrollment.StudyPlaceId, id)
}
