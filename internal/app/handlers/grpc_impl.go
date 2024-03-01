package handlers

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/grpc"
	"github.com/stdyum/api-common/proto/impl/types_registry"
)

func (h *gRPC) GetGroupsByIds(ctx context.Context, list *types_registry.IdList) (*types_registry.Groups, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, list.Token, list.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(list.Ids))
	for i, id := range list.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, grpc.ConvertError(err)
		}

		ids[i] = uid
	}

	groups, err := h.controller.GetGroupsByIds(ctx, enrollmentUser.Enrollment, ids)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	groupsRpc := make(map[string]*types_registry.Group, len(groups))
	for _, group := range groups {
		id := group.ID.String()
		groupsRpc[id] = &types_registry.Group{
			Id:   id,
			Name: group.Name,
		}
	}

	return &types_registry.Groups{List: groupsRpc}, nil
}

func (h *gRPC) GetRoomsByIds(ctx context.Context, list *types_registry.IdList) (*types_registry.Rooms, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, list.Token, list.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(list.Ids))
	for i, id := range list.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, grpc.ConvertError(err)
		}

		ids[i] = uid
	}

	rooms, err := h.controller.GetRoomsByIds(ctx, enrollmentUser.Enrollment, ids)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	roomsRpc := make(map[string]*types_registry.Room, len(rooms))
	for _, room := range rooms {
		id := room.ID.String()
		roomsRpc[id] = &types_registry.Room{
			Id:   id,
			Name: room.Name,
		}
	}

	return &types_registry.Rooms{List: roomsRpc}, nil
}

func (h *gRPC) GetStudentsByIds(ctx context.Context, list *types_registry.IdList) (*types_registry.Students, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, list.Token, list.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(list.Ids))
	for i, id := range list.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, grpc.ConvertError(err)
		}

		ids[i] = uid
	}

	students, err := h.controller.GetStudentsByIds(ctx, enrollmentUser.Enrollment, ids)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	studentsRpc := make(map[string]*types_registry.Student, len(students))
	for _, student := range students {
		id := student.ID.String()
		studentsRpc[id] = &types_registry.Student{
			Id:   id,
			Name: student.Name,
		}
	}

	return &types_registry.Students{List: studentsRpc}, nil
}

func (h *gRPC) GetSubjectsByIds(ctx context.Context, list *types_registry.IdList) (*types_registry.Subjects, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, list.Token, list.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(list.Ids))
	for i, id := range list.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, grpc.ConvertError(err)
		}

		ids[i] = uid
	}

	subjects, err := h.controller.GetSubjectsByIds(ctx, enrollmentUser.Enrollment, ids)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	subjectsRpc := make(map[string]*types_registry.Subject, len(subjects))
	for _, subject := range subjects {
		id := subject.ID.String()
		subjectsRpc[id] = &types_registry.Subject{
			Id:   id,
			Name: subject.Name,
		}
	}

	return &types_registry.Subjects{List: subjectsRpc}, nil
}

func (h *gRPC) GetTeachersByIds(ctx context.Context, list *types_registry.IdList) (*types_registry.Teachers, error) {
	enrollmentUser, err := grpc.EnrollmentAuth(ctx, list.Token, list.StudyPlaceId)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, len(list.Ids))
	for i, id := range list.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return nil, grpc.ConvertError(err)
		}

		ids[i] = uid
	}

	teachers, err := h.controller.GetTeachersByIds(ctx, enrollmentUser.Enrollment, ids)
	if err != nil {
		return nil, grpc.ConvertError(err)
	}

	teachersRpc := make(map[string]*types_registry.Teacher, len(teachers))
	for _, teacher := range teachers {
		id := teacher.ID.String()
		teachersRpc[id] = &types_registry.Teacher{
			Id:   id,
			Name: teacher.Name,
		}
	}

	return &types_registry.Teachers{List: teachersRpc}, nil
}
