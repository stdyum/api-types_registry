package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/mock"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-types-registry/internal"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func main() {
	_, ctrl, err := internal.Configure()
	if err != nil {
		log.Fatal(err)
	}

	mocker := mock.Mock{
		Config: mock.Config{
			MaxSingleInsert: 50,
		},
		Item: mock.DataItem[any, models.Enrollment]{
			GenerateItems: func(ctx context.Context, data *any) ([]models.Enrollment, error) {
				return []models.Enrollment{
					{
						StudyPlaceId: uuid.MustParse("7115bc03-143f-4acd-912d-02ffa9f14462"),
						Permissions:  []models.Permission{models.PermissionAdmin},
					},
				}, nil
			},
			Nested: []mock.RawItem{
				mock.DataItemNested[any, dto.CreateGroupEntryRequestDTO, models.Enrollment]{
					Config: mock.ConfigItem{
						Amount: 10,
					},
					Generate: func(ctx context.Context, i int, data *any, previous models.Enrollment) (dto.CreateGroupEntryRequestDTO, error) {
						return dto.CreateGroupEntryRequestDTO{Name: "Group" + strconv.Itoa(i)}, nil
					},
					Insert: func(ctx context.Context, i int, items []dto.CreateGroupEntryRequestDTO, previous models.Enrollment) error {
						_, err = ctrl.CreateGroups(ctx, previous, dto.CreateGroupsRequestDTO{List: items})
						return err
					},
				}.Build(),
				mock.DataItemNested[any, dto.CreateRoomEntryRequestDTO, models.Enrollment]{
					Config: mock.ConfigItem{
						Amount: 50,
					},
					Generate: func(ctx context.Context, i int, data *any, previous models.Enrollment) (dto.CreateRoomEntryRequestDTO, error) {
						return dto.CreateRoomEntryRequestDTO{Name: "Room" + strconv.Itoa(i)}, nil
					},
					Insert: func(ctx context.Context, i int, items []dto.CreateRoomEntryRequestDTO, previous models.Enrollment) error {
						_, err = ctrl.CreateRooms(ctx, previous, dto.CreateRoomsRequestDTO{List: items})
						return err
					},
				}.Build(),
				mock.DataItemNested[any, dto.CreateSubjectEntryRequestDTO, models.Enrollment]{
					Config: mock.ConfigItem{
						Amount: 50,
					},
					Generate: func(ctx context.Context, i int, data *any, previous models.Enrollment) (dto.CreateSubjectEntryRequestDTO, error) {
						return dto.CreateSubjectEntryRequestDTO{Name: "Room" + strconv.Itoa(i)}, nil
					},
					Insert: func(ctx context.Context, i int, items []dto.CreateSubjectEntryRequestDTO, previous models.Enrollment) error {
						_, err = ctrl.CreateSubjects(ctx, previous, dto.CreateSubjectsRequestDTO{List: items})
						return err
					},
				}.Build(),
				mock.DataItemNested[any, dto.CreateTeacherEntryRequestDTO, models.Enrollment]{
					Config: mock.ConfigItem{
						Amount: 50,
					},
					Generate: func(ctx context.Context, i int, data *any, previous models.Enrollment) (dto.CreateTeacherEntryRequestDTO, error) {
						return dto.CreateTeacherEntryRequestDTO{Name: "Room" + strconv.Itoa(i)}, nil
					},
					Insert: func(ctx context.Context, i int, items []dto.CreateTeacherEntryRequestDTO, previous models.Enrollment) error {
						_, err = ctrl.CreateTeachers(ctx, previous, dto.CreateTeachersRequestDTO{List: items})
						return err
					},
				}.Build(),
			},
		}.Build(),
	}

	if err = mocker.Mock(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
