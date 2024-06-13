package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/databases/query_builder"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (r *repository) GetStudentsInGroup(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID) ([]entities.Student, error) {
	scanner, err := r.database.QueryContext(ctx, `
SELECT students.id, students.study_place_id, students.name, students.created_at, students.updated_at
FROM student_groups
         INNER JOIN students ON students.id = student_groups.student_id
WHERE student_groups.study_place_id = $1
  AND student_groups.group_id = $2
`, studyPlaceId, groupId,
	)
	return databases.ScanArrayErr(scanner, r.scanStudent, err)
}

func (r *repository) GetStudentsInGroupsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.AggregatedStudentGroup, int, error) {
	paginationQuery.SetField("student_groups.created_at")
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		`
SELECT students.study_place_id, students.id, students.name, groups.id, groups.name, student_groups.created_at, student_groups.updated_at
FROM student_groups
         INNER JOIN students ON students.id = student_groups.student_id
         INNER JOIN groups ON groups.id = student_groups.group_id
WHERE student_groups.study_place_id = $1
`,
		"SELECT count(*) FROM student_groups WHERE study_place_id = $1",
		paginationQuery,
		studyPlaceId,
	)

	return databases.ScanPaginationErr(result, r.scanAggregatedStudentGroup, total, err)
}

func (r *repository) GetStudentGroups(ctx context.Context, studyPlaceId uuid.UUID, studentId uuid.UUID) ([]entities.Group, error) {
	scanner, err := r.database.QueryContext(ctx, `
SELECT groups.id, groups.study_place_id, groups.name, groups.created_at, groups.updated_at
FROM student_groups
         INNER JOIN groups ON groups.id = student_groups.group_id
WHERE student_groups.study_place_id = $1
  AND student_groups.student_id = $2
`, studyPlaceId, studentId,
	)
	return databases.ScanArrayErr(scanner, r.scanGroup, err)
}

func (r *repository) AddStudentsToGroup(ctx context.Context, studyPlaceId uuid.UUID, studentIds []uuid.UUID, groupId uuid.UUID) error {
	builder := query_builder.NewQueryBuilder(`
INSERT INTO student_groups (study_place_id, student_id, group_id) VALUES 
	`)

	for _, id := range studentIds {
		builder.Append("($1, $2, $3),", studyPlaceId, id, groupId)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) RemoveStudentFromGroup(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID, studentId uuid.UUID) error {
	result, err := r.database.ExecContext(ctx, `
DELETE FROM student_groups WHERE study_place_id = $1 AND student_id = $2 AND group_id = $3
	`, studyPlaceId, studentId, groupId)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) GetTeacherTuitionGroups(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID) ([]entities.Group, error) {
	scanner, err := r.database.QueryContext(ctx, `
SELECT groups.id, groups.study_place_id, groups.name, groups.created_at, groups.updated_at
FROM tutor_groups
         INNER JOIN groups ON groups.id = tutor_groups.group_id
WHERE tutor_groups.study_place_id = $1
  AND tutor_groups.group_id = $2
`, studyPlaceId, groupId,
	)
	return databases.ScanArrayErr(scanner, r.scanGroup, err)
}

func (r *repository) AddTutorToGroups(ctx context.Context, studyPlaceId uuid.UUID, groupIds []uuid.UUID, teacherId uuid.UUID) error {
	builder := query_builder.NewQueryBuilder(`
INSERT INTO tutor_groups (study_place_id, teacher_id, group_id) VALUES 
	`)

	for _, id := range groupIds {
		builder.Append("($1, $2, $3),", studyPlaceId, teacherId, id)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) RemoveGroupTutor(ctx context.Context, studyPlaceId uuid.UUID, groupId uuid.UUID, teacherId uuid.UUID) error {
	result, err := r.database.ExecContext(ctx, `
DELETE FROM tutor_groups WHERE study_place_id = $1 AND teacher_id = $2 AND group_id = $3
	`, studyPlaceId, teacherId, groupId)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) GetGroupIdsWithStudents(ctx context.Context, studyPlaceId uuid.UUID) ([]uuid.UUID, error) {
	scanner, err := r.database.QueryContext(ctx, `
SELECT DISTINCT group_id FROM student_groups WHERE study_place_id = $1
`, studyPlaceId,
	)

	return databases.ScanArrayErr(scanner, r.scanUUID, err)
}
