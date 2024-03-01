package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-types-registry/internal/app/dto"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (c *controller) GetStudentsInGroup(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.StudentItemResponseDTO, error) {
	students, err := c.repository.GetStudentsInGroup(ctx, enrollment.StudyPlaceId, groupId)
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

func (c *controller) AddStudentsToGroup(ctx context.Context, enrollment models.Enrollment, request dto.AddStudentsToGroupRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.AddStudentsToGroup(ctx, enrollment.StudyPlaceId, request.StudentIds, request.GroupId)
}

func (c *controller) RemoveStudentFromGroup(ctx context.Context, enrollment models.Enrollment, request dto.RemoveStudentFromGroupDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.RemoveStudentFromGroup(ctx, enrollment.StudyPlaceId, request.GroupId, request.StudentId)
}

func (c *controller) GetTeacherTuitionGroups(ctx context.Context, enrollment models.Enrollment, groupId uuid.UUID) ([]dto.GroupItemResponseDTO, error) {
	groups, err := c.repository.GetTeacherTuitionGroups(ctx, enrollment.StudyPlaceId, groupId)
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

func (c *controller) AddTutorToGroups(ctx context.Context, enrollment models.Enrollment, request dto.AddTutorToGroupsRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.AddTutorToGroups(ctx, enrollment.StudyPlaceId, request.GroupIds, request.TeacherId)
}

func (c *controller) RemoveGroupTutor(ctx context.Context, enrollment models.Enrollment, request dto.RemoveGroupTutorRequestDTO) error {
	if err := enrollment.Permissions.Assert(models.PermissionRegistry); err != nil {
		return err
	}

	return c.repository.RemoveGroupTutor(ctx, enrollment.StudyPlaceId, request.GroupId, request.TeacherId)
}
