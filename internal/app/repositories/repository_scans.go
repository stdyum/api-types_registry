package repositories

import (
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (r *repository) scanRoom(row databases.Scan) (room entities.Room, err error) {
	err = row.Scan(
		&room.ID,
		&room.StudyPlaceId,
		&room.Name,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	return
}

func (r *repository) scanGroup(row databases.Scan) (group entities.Group, err error) {
	err = row.Scan(
		&group.ID,
		&group.StudyPlaceId,
		&group.Name,
		&group.CreatedAt,
		&group.UpdatedAt,
	)
	return
}

func (r *repository) scanStudent(row databases.Scan) (student entities.Student, err error) {
	err = row.Scan(
		&student.ID,
		&student.StudyPlaceId,
		&student.Name,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	return
}

func (r *repository) scanSubject(row databases.Scan) (subject entities.Subject, err error) {
	err = row.Scan(
		&subject.ID,
		&subject.StudyPlaceId,
		&subject.Name,
		&subject.CreatedAt,
		&subject.UpdatedAt,
	)
	return
}

func (r *repository) scanTeacher(row databases.Scan) (teacher entities.Teacher, err error) {
	err = row.Scan(
		&teacher.ID,
		&teacher.StudyPlaceId,
		&teacher.Name,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)
	return
}
