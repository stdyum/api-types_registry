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

func (r *repository) GetGroupsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Group, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, study_place_id, name, created_at, updated_at FROM groups WHERE study_place_id = $1",
		"SELECT count(*) FROM groups WHERE study_place_id = $1",
		paginationQuery,
		studyPlaceId,
	)
	return databases.ScanPaginationErr(result, r.scanGroup, total, err)
}

func (r *repository) GetGroupsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Group, error) {
	result, err := r.database.QueryContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM groups WHERE id = ANY($1) AND study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.ScanArrayErr(result, r.scanGroup, err)
}

func (r *repository) GetGroupById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Group, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM groups WHERE id = $1 AND study_place_id = $2",
		id, studyPlaceId,
	)

	return r.scanGroup(row)
}

func (r *repository) CreateGroups(ctx context.Context, groups []entities.Group) error {
	builder := query_builder.NewQueryBuilder("INSERT INTO groups (id, study_place_id, name) VALUES ")
	for _, group := range groups {
		builder.Append("($1, $2, $3),", group.ID, group.StudyPlaceId, group.Name)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) UpdateGroup(ctx context.Context, group entities.Group) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE groups SET name = $1 WHERE id = $2 and study_place_id = $3",
		group.Name, group.ID, group.StudyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteGroupsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM groups WHERE id = ANY($1) and study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}
