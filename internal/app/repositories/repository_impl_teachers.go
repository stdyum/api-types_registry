package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stdyum/api-common/databases"
	"github.com/stdyum/api-common/databases/pagination"
	"github.com/stdyum/api-common/databases/query_builder"
	"github.com/stdyum/api-types-registry/internal/app/entities"
)

func (r *repository) GetTeachersPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Teacher, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, study_place_id, name, created_at, updated_at FROM teachers WHERE study_place_id = $1",
		"SELECT count(*) FROM teachers WHERE study_place_id = $1",
		paginationQuery,
		[]string{"name"},
		studyPlaceId,
	)
	return databases.ScanPaginationErr(result, r.scanTeacher, total, err)
}

func (r *repository) GetTeachersByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Teacher, error) {
	result, err := r.database.QueryContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM teachers WHERE id = ANY($1) AND study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.ScanArrayErr(result, r.scanTeacher, err)
}

func (r *repository) GetTeacherById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Teacher, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM teachers WHERE id = $1 AND study_place_id = $2",
		id, studyPlaceId,
	)

	return r.scanTeacher(row)
}

func (r *repository) CreateTeachers(ctx context.Context, teachers []entities.Teacher) error {
	builder := query_builder.NewQueryBuilder("INSERT INTO teachers (id, study_place_id, name) VALUES ")
	for _, teacher := range teachers {
		builder.Append("($1, $2, $3),", teacher.ID, teacher.StudyPlaceId, teacher.Name)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) UpdateTeacher(ctx context.Context, subject entities.Teacher) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE teachers SET name = $1 WHERE id = $2 and study_place_id = $3",
		subject.Name, subject.ID, subject.StudyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteTeachersByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM teachers WHERE id = ANY($1) and study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}
