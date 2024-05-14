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

func (c *controller) GetSubjectsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.SubjectsResponseDTO, error) {
	groups, amount, err := c.repository.GetSubjectsPaginated(ctx, studyPlaceId, paginationQuery)
	if err != nil {
		return dto.SubjectsResponseDTO{}, err
	}

	paginationResult := pagination.FromArrayAndAmount(groups, amount, paginationQuery,
		func(el entities.Subject) dto.SubjectItemResponseDTO {
			return dto.SubjectItemResponseDTO{
				ID:           el.ID,
				StudyPlaceId: el.StudyPlaceId,
				Name:         el.Name,
			}
		},
	)

	return dto.SubjectsResponseDTO(paginationResult), nil
}

func (c *controller) GetSubjectsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.SubjectItemResponseDTO, error) {
	subjects, err := c.repository.GetSubjectsByIds(ctx, enrollment.StudyPlaceId, ids)
	if err != nil {
		return nil, err
	}

	response := uslices.MapFunc(subjects, func(item entities.Subject) dto.SubjectItemResponseDTO {
		return dto.SubjectItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	})

	return response, nil
}

func (c *controller) GetSubjectById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.SubjectItemResponseDTO, error) {
	subject, err := c.repository.GetSubjectById(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.SubjectItemResponseDTO{}, err
	}

	return dto.SubjectItemResponseDTO{
		ID:           subject.ID,
		StudyPlaceId: subject.StudyPlaceId,
		Name:         subject.Name,
	}, nil
}

func (c *controller) CreateSubjects(ctx context.Context, enrollment models.Enrollment, request dto.CreateSubjectsRequestDTO) ([]dto.SubjectItemResponseDTO, error) {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return nil, err
	}

	subjects := uslices.MapFunc(request.List, func(item dto.CreateSubjectEntryRequestDTO) entities.Subject {
		return entities.Subject{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Name:         item.Name,
		}
	})

	if err := c.repository.CreateSubjects(ctx, subjects); err != nil {
		return nil, err
	}

	return uslices.MapFunc(subjects, func(item entities.Subject) dto.SubjectItemResponseDTO {
		return dto.SubjectItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	}), nil
}

func (c *controller) UpdateSubject(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateSubjectRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	subject := entities.Subject{
		ID:           id,
		StudyPlaceId: enrollment.StudyPlaceId,
		Name:         request.Name,
	}

	return c.repository.UpdateSubject(ctx, subject)
}

func (c *controller) DeleteSubjectsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteSubjectsByIdsRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.DeleteSubjectsByIds(ctx, enrollment.StudyPlaceId, request.IDs)
}
