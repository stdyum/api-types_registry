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

func (r *repository) GetSubjectsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Subject, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, study_place_id, name, created_at, updated_at FROM subjects WHERE study_place_id = $1",
		"SELECT count(*) FROM subjects",
		paginationQuery,
		studyPlaceId,
	)
	return databases.ScanPaginationErr(result, r.scanSubject, total, err)
}

func (r *repository) GetSubjectsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Subject, error) {
	result, err := r.database.QueryContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM subjects WHERE id = ANY($1) AND study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.ScanArrayErr(result, r.scanSubject, err)
}

func (r *repository) GetSubjectById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Subject, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM subjects WHERE id = $1 AND study_place_id = $2",
		id, studyPlaceId,
	)

	return r.scanSubject(row)
}

func (r *repository) CreateSubjects(ctx context.Context, subjects []entities.Subject) error {
	builder := query_builder.NewQueryBuilder("INSERT INTO subjects (id, study_place_id, name) VALUES ")
	for _, subject := range subjects {
		builder.Append("($1, $2, $3),", subject.ID, subject.StudyPlaceId, subject.Name)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) UpdateSubject(ctx context.Context, subject entities.Subject) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE subjects SET name = $1 WHERE id = $2 and study_place_id = $3",
		subject.Name, subject.ID, subject.StudyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteSubjectsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM subjects WHERE id = ANY($1) and study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}
