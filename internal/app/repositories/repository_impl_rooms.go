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

func (r *repository) GetRoomsPaginated(ctx context.Context, studyPlaceId uuid.UUID, paginationQuery pagination.Query) ([]entities.Room, int, error) {
	result, total, err := pagination.QueryPaginationContext(
		ctx, r.database,
		"SELECT id, study_place_id, name, created_at, updated_at FROM rooms WHERE study_place_id = $1",
		"SELECT count(*) FROM rooms WHERE study_place_id = $1",
		paginationQuery,
		[]string{"name"},
		studyPlaceId,
	)
	return databases.ScanPaginationErr(result, r.scanRoom, total, err)

}

func (r *repository) GetRoomsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) ([]entities.Room, error) {
	result, err := r.database.QueryContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM rooms WHERE id = ANY($1) AND study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.ScanArrayErr(result, r.scanRoom, err)
}

func (r *repository) GetRoomById(ctx context.Context, studyPlaceId uuid.UUID, id uuid.UUID) (entities.Room, error) {
	row := r.database.QueryRowContext(ctx,
		"SELECT id, study_place_id, name, created_at, updated_at FROM rooms WHERE id = $1 AND study_place_id = $2",
		id, studyPlaceId,
	)

	return r.scanRoom(row)
}

func (r *repository) CreateRooms(ctx context.Context, rooms []entities.Room) error {
	builder := query_builder.NewQueryBuilder("INSERT INTO rooms (id, study_place_id, name) VALUES ")
	for _, room := range rooms {
		builder.Append("($1, $2, $3),", room.ID, room.StudyPlaceId, room.Name)
	}
	builder.RemoveLast(1)

	query, args := builder.Build()
	_, err := r.database.ExecContext(ctx, query, args...)
	return err
}

func (r *repository) UpdateRoom(ctx context.Context, room entities.Room) error {
	result, err := r.database.ExecContext(ctx,
		"UPDATE rooms SET name = $1 WHERE id = $2 and study_place_id = $3",
		room.Name, room.ID, room.StudyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}

func (r *repository) DeleteRoomsByIds(ctx context.Context, studyPlaceId uuid.UUID, ids []uuid.UUID) error {
	result, err := r.database.ExecContext(ctx,
		"DELETE FROM rooms WHERE id = ANY($1) and study_place_id = $2",
		pq.Array(ids), studyPlaceId,
	)
	return databases.AssertRowAffectedErr(result, err)
}
