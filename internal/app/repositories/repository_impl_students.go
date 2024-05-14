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

func (r *repository) GetStudentsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Student, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, study_place_id, name, created_at, updated_at FROM students WHERE study_place_id = $1",
		"SELECT count(*) FROM students WHERE study_place_id = $1",
		paginationQuery,
		studyPlaceId,
	)
	return databases.ScanPaginationErr(result, r.scanStudent, total, err)

}

func (r *repository) GetStudentsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Student, error) {
	result, err := r.database.QueryContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM students WHERE id = ANY($1) AND study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.ScanArrayErr(result, r.scanStudent, err)
}

func (r *repository) GetStudentById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Student, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM students WHERE id = $1 AND study_place_id = $2",
		id, studyPlaceId,
	)

	return r.scanStudent(row)
}

func (r *repository) CreateStudents(ctx context.Context, students []entities.Student) error {
	builder := query_builder.NewQueryBuilder("INSERT INTO students (id, study_place_id, name) VALUES ")
	for _, student := range students {
		builder.Append("($1, $2, $3),", student.ID, student.StudyPlaceId, student.Name)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) UpdateStudent(ctx context.Context, student entities.Student) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE students SET name = $1 WHERE id = $2 and study_place_id = $3",
		student.Name, student.ID, student.StudyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteStudentsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM students WHERE id = ANY($1) and study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}
