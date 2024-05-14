package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-types-registry/internal/app/dto"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (c *controller) GetTypesById(ctx context.Context, enrollment models.Enrollment, ids dto.GetTypesByIdRequestDTO) (response dto.TypesModelsResponseDTO, err error) {
	if len(ids.GroupsIds) != 0 {
		groups, err := c.repository.GetGroupsByIds(ctx, enrollment.StudyPlaceId, ids.GroupsIds)
		if err != nil {
			return dto.TypesModelsResponseDTO{}, err
		}

		response.GroupsIds = uslices.ToMapFunc(groups, func(item entities.Group) (string, dto.TypesModelsGroupResponseDTO) {
			return item.ID.String(), dto.TypesModelsGroupResponseDTO{
				ID:   item.ID.String(),
				Name: item.Name,
			}
		})
	}

	if len(ids.RoomsIds) != 0 {
		rooms, err := c.repository.GetRoomsByIds(ctx, enrollment.StudyPlaceId, ids.RoomsIds)
		if err != nil {
			return dto.TypesModelsResponseDTO{}, err
		}

		response.RoomsIds = uslices.ToMapFunc(rooms, func(item entities.Room) (string, dto.TypesModelsRoomResponseDTO) {
			return item.ID.String(), dto.TypesModelsRoomResponseDTO{
				ID:   item.ID.String(),
				Name: item.Name,
			}
		})
	}

	if len(ids.StudentIds) != 0 {
		subjects, err := c.repository.GetStudentsByIds(ctx, enrollment.StudyPlaceId, ids.SubjectsIds)
		if err != nil {
			return dto.TypesModelsResponseDTO{}, err
		}

		response.StudentsIds = uslices.ToMapFunc(subjects, func(item entities.Student) (string, dto.TypesModelsStudentResponseDTO) {
			return item.ID.String(), dto.TypesModelsStudentResponseDTO{
				ID:   item.ID.String(),
				Name: item.Name,
			}
		})
	}

	if len(ids.SubjectsIds) != 0 {
		subjects, err := c.repository.GetSubjectsByIds(ctx, enrollment.StudyPlaceId, ids.SubjectsIds)
		if err != nil {
			return dto.TypesModelsResponseDTO{}, err
		}

		response.SubjectsIds = uslices.ToMapFunc(subjects, func(item entities.Subject) (string, dto.TypesModelsSubjectResponseDTO) {
			return item.ID.String(), dto.TypesModelsSubjectResponseDTO{
				ID:   item.ID.String(),
				Name: item.Name,
			}
		})
	}

	if len(ids.TeachersIds) != 0 {
		teachers, err := c.repository.GetTeachersByIds(ctx, enrollment.StudyPlaceId, ids.TeachersIds)
		if err != nil {
			return dto.TypesModelsResponseDTO{}, err
		}

		response.TeachersIds = uslices.ToMapFunc(teachers, func(item entities.Teacher) (string, dto.TypesModelsTeacherResponseDTO) {
			return item.ID.String(), dto.TypesModelsTeacherResponseDTO{
				ID:   item.ID.String(),
				Name: item.Name,
			}
		})
	}

	return
}

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
