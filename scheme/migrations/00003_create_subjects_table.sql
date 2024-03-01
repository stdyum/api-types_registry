-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS subjects
(
    id   uuid primary key not null default uuid_generate_v4(),
    study_place_id uuid             not null,
    name           varchar          not null
);

CALL register_updated_at_created_at_columns('subjects');

CREATE INDEX IF NOT EXISTS subjects_created_at_idx ON subjects USING hash (created_at);

-- +goose StatementEnd
