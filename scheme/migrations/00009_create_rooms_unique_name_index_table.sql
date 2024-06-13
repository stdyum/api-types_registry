-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX idx_rooms_name_unique ON rooms (study_place_id, name)

-- +goose StatementEnd
