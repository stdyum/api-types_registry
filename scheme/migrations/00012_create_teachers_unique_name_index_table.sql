-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX idx_teachers_name_unique ON teachers (study_place_id, name)

-- +goose StatementEnd
