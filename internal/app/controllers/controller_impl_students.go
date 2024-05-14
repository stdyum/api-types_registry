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

func (c *controller) GetStudentsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) (dto.StudentsResponseDTO, error) {
	students, amount, err := c.repository.GetStudentsPaginated(ctx, studyPlaceId, paginationQuery)
	if err != nil {
		return dto.StudentsResponseDTO{}, err
	}

	paginationResult := pagination.FromArrayAndAmount(students, amount, paginationQuery,
		func(el entities.Student) dto.StudentItemResponseDTO {
			return dto.StudentItemResponseDTO{
				ID:           el.ID,
				StudyPlaceId: el.StudyPlaceId,
				Name:         el.Name,
			}
		},
	)

	return dto.StudentsResponseDTO(paginationResult), nil
}

func (c *controller) GetStudentsByIds(ctx context.Context, enrollment models.Enrollment, ids []uuid.UUID) ([]dto.StudentItemResponseDTO, error) {
	students, err := c.repository.GetStudentsByIds(ctx, enrollment.StudyPlaceId, ids)
	if err != nil {
		return nil, err
	}

	response := uslices.MapFunc(students, func(item entities.Student) dto.StudentItemResponseDTO {
		return dto.StudentItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	})

	return response, nil
}

func (c *controller) GetStudentById(ctx context.Context, enrollment models.Enrollment, id uuid.UUID) (dto.StudentItemResponseDTO, error) {
	student, err := c.repository.GetStudentById(ctx, enrollment.StudyPlaceId, id)
	if err != nil {
		return dto.StudentItemResponseDTO{}, err
	}

	return dto.StudentItemResponseDTO{
		ID:           student.ID,
		StudyPlaceId: student.StudyPlaceId,
		Name:         student.Name,
	}, nil
}

func (c *controller) CreateStudents(ctx context.Context, enrollment models.Enrollment, request dto.CreateStudentsRequestDTO) ([]dto.StudentItemResponseDTO, error) {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return nil, err
	}

	students := uslices.MapFunc(request.List, func(item dto.CreateStudentEntryRequestDTO) entities.Student {
		return entities.Student{
			ID:           uuid.New(),
			StudyPlaceId: enrollment.StudyPlaceId,
			Name:         item.Name,
		}
	})

	if err := c.repository.CreateStudents(ctx, students); err != nil {
		return nil, err
	}

	return uslices.MapFunc(students, func(item entities.Student) dto.StudentItemResponseDTO {
		return dto.StudentItemResponseDTO{
			ID:           item.ID,
			StudyPlaceId: item.StudyPlaceId,
			Name:         item.Name,
		}
	}), nil
}

func (c *controller) UpdateStudent(ctx context.Context, enrollment models.Enrollment, id uuid.UUID, request dto.UpdateStudentRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	student := entities.Student{
		ID:           id,
		StudyPlaceId: enrollment.StudyPlaceId,
		Name:         request.Name,
	}

	return c.repository.UpdateStudent(ctx, student)
}

func (c *controller) DeleteStudentsByIds(ctx context.Context, enrollment models.Enrollment, request dto.DeleteStudentsByIdsRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.DeleteStudentsByIds(ctx, enrollment.StudyPlaceId, request.IDs)
}
