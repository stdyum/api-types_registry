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

func (c *controller) GetTeachersPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.TeachersResponseDTO, error) {
	teachers, amount, err := c.repository.GetTeachersPaginated(ctx, studyPlaceId, paginationQuery)
	if err != nil {
		return dto.TeachersResponseDTO{}, err
	}

	paginationResult := pagination.FromArrayAndAmount(teachers, amount, paginationQuery,
		func(el entities.Teacher) dto.TeacherItemResponseDTO {
			return dto.TeacherItemResponseDTO{
				ID:           el.ID,
				StudyPlaceId: el.StudyPlaceId,
				Name:         el.Name,
			}
		},
	)

	return dto.TeachersResponseDTO(paginationResult), nil
}

func (c *controller) GetTeachersByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.TeacherItemResponseDTO, error) {
	teachers, err := c.repository.GetTeachersByIds(ctx, enrollment.StudyPlaceId, ids)
	if err != nil {
		return nil, err
	}

	response := uslices.MapFunc(teachers, func(item entities.Teacher) dto.TeacherItemResponseDTO {
		return dto.TeacherItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	})

	return response, nil
}

func (c *controller) GetTeacherById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.TeacherItemResponseDTO, error) {
	teacher, err := c.repository.GetTeacherById(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.TeacherItemResponseDTO{}, err
	}

	return dto.TeacherItemResponseDTO{
		ID:           teacher.ID,
		StudyPlaceId: teacher.StudyPlaceId,
		Name:         teacher.Name,
	}, nil
}

func (c *controller) CreateTeachers(ctx context.Context, enrollment models.Enrollment, request dto.CreateTeachersRequestDTO) ([]dto.TeacherItemResponseDTO, error) {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return nil, err
	}

	teachers := uslices.MapFunc(request.List, func(item dto.CreateTeacherEntryRequestDTO) entities.Teacher {
		return entities.Teacher{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Name:         item.Name,
		}
	})

	if err := c.repository.CreateTeachers(ctx, teachers); err != nil {
		return nil, err
	}

	return uslices.MapFunc(teachers, func(item entities.Teacher) dto.TeacherItemResponseDTO {
		return dto.TeacherItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	}), nil
}

func (c *controller) UpdateTeacher(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateTeacherRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	teacher := entities.Teacher{
		ID:           id,
		StudyPlaceId: enrollment.StudyPlaceId,
		Name:         request.Name,
	}

	return c.repository.UpdateTeacher(ctx, teacher)
}

func (c *controller) DeleteTeachersByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteTeachersByIdsRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.DeleteTeachersByIds(ctx, enrollment.StudyPlaceId, request.IDs)
}
