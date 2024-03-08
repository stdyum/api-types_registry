package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-common/proto/impl/types_registry"
	"github.com/stdyum/api-common/umaps"
	"github.com/stdyum/api-common/uslices"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func (h *gRPC) GetTypesByIds(ctx context.Context, ids *types_registry.TypesIds) (*types_registry.TypesModels, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, ids.Token, ids.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	var typesIds models.TypesIds

	typesIds.GroupsIds, err = uslices.MapFuncErr(ids.GroupsIds, func(item string) (uuid.UUID, error) {
		return uuid.Parse(item)
	})
	if err != nil {
		return nil, err
	}

	typesIds.RoomsIds, err = uslices.MapFuncErr(ids.RoomsIds, func(item string) (uuid.UUID, error) {
		return uuid.Parse(item)
	})
	if err != nil {
		return nil, err
	}

	typesIds.SubjectsIds, err = uslices.MapFuncErr(ids.SubjectsIds, func(item string) (uuid.UUID, error) {
		return uuid.Parse(item)
	})
	if err != nil {
		return nil, err
	}

	typesIds.TeachersIds, err = uslices.MapFuncErr(ids.TeachersIds, func(item string) (uuid.UUID, error) {
		return uuid.Parse(item)
	})
	if err != nil {
		return nil, err
	}

	typesModel, err := h.controller.GetTypesById(ctx, enrollmentUser.Enrollment, dto.GetTypesByIdRequestDTO(typesIds))
	if err != nil {
		return nil, err
	}

	return &types_registry.TypesModels{
		Groups: umaps.MapFunc(typesModel.GroupsIds, func(key string, value dto.TypesModelsGroupResponseDTO) (string, *types_registry.Group) {
			return key, &types_registry.Group{
				Id:   value.ID,
				Name: value.Name,
			}
		}),
		Rooms: umaps.MapFunc(typesModel.RoomsIds, func(key string, value dto.TypesModelsRoomResponseDTO) (string, *types_registry.Room) {
			return key, &types_registry.Room{
				Id:   value.ID,
				Name: value.Name,
			}
		}),
		Subjects: umaps.MapFunc(typesModel.SubjectsIds, func(key string, value dto.TypesModelsSubjectResponseDTO) (string, *types_registry.Subject) {
			return key, &types_registry.Subject{
				Id:   value.ID,
				Name: value.Name,
			}
		}),
		Teachers: umaps.MapFunc(typesModel.TeachersIds, func(key string, value dto.TypesModelsTeacherResponseDTO) (string, *types_registry.Teacher) {
			return key, &types_registry.Teacher{
				Id:   value.ID,
				Name: value.Name,
			}
		}),
	}, nil
}
