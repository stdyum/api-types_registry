package main

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/models"
	"github.com/stdyum/api-types-registry/internal"
	"github.com/stdyum/api-types-registry/internal/app/controllers"
	"github.com/stdyum/api-types-registry/internal/app/dto"
)

func main() {
	_, ctrl, err := internal.Configure()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	enrollments := []models.Enrollment{
		{
			StudyPlaceId: uuid.MustParse("22dfd1be-eb72-4e1e-8659-cf4b8d95e90d"),
			Permissions:  []models.Permission{models.PermissionAdmin},
		},
	}

	for _, enrollment := range enrollments {
		mockEnrollment(ctx, ctrl, enrollment)
	}

}

func mockEnrollment(ctx context.Context, ctrl controllers.Controller, enrollment models.Enrollment) {
	var (
		groupsAmount   = 10
		roomsAmount    = 50
		subjectsAmount = 40
		teachersAmount = 30
	)

	groupsRequest := dto.CreateGroupsRequestDTO{
		List: make([]dto.CreateGroupEntryRequestDTO, groupsAmount),
	}

	for i := 0; i < groupsAmount; i++ {
		groupsRequest.List[i] = dto.CreateGroupEntryRequestDTO{
			Name: "Group" + strconv.Itoa(i),
		}
	}

	roomsRequest := dto.CreateRoomsRequestDTO{
		List: make([]dto.CreateRoomEntryRequestDTO, roomsAmount),
	}

	for i := 0; i < roomsAmount; i++ {
		roomsRequest.List[i] = dto.CreateRoomEntryRequestDTO{
			Name: "Room" + strconv.Itoa(i),
		}
	}

	subjectsRequest := dto.CreateSubjectsRequestDTO{
		List: make([]dto.CreateSubjectEntryRequestDTO, subjectsAmount),
	}

	for i := 0; i < subjectsAmount; i++ {
		subjectsRequest.List[i] = dto.CreateSubjectEntryRequestDTO{
			Name: "Subject" + strconv.Itoa(i),
		}
	}

	teachersRequest := dto.CreateTeachersRequestDTO{
		List: make([]dto.CreateTeacherEntryRequestDTO, teachersAmount),
	}

	for i := 0; i < teachersAmount; i++ {
		teachersRequest.List[i] = dto.CreateTeacherEntryRequestDTO{
			Name: "Teacher" + strconv.Itoa(i),
		}
	}

	_, err := ctrl.CreateGroups(ctx, enrollment, groupsRequest)
	if err != nil {
		log.Println(err)
	}

	_, err = ctrl.CreateRooms(ctx, enrollment, roomsRequest)
	if err != nil {
		log.Println(err)
	}

	_, err = ctrl.CreateSubjects(ctx, enrollment, subjectsRequest)
	if err != nil {
		log.Println(err)
	}

	_, err = ctrl.CreateTeachers(ctx, enrollment, teachersRequest)
	if err != nil {
		log.Println(err)
	}
}
