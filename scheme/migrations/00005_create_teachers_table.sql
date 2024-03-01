-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS teachers
(
    id   uuid primary key not null default uuid_generate_v4(),
    study_place_id uuid             not null,
    name           varchar          not null
);

CALL register_updated_at_created_at_columns('teachers');

CREATE INDEX IF NOT EXISTS teachers_created_at_idx ON teachers USING hash (created_at);

-- +goose StatementEnd
