-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX idx_students_name_unique ON students (study_place_id, name)

-- +goose StatementEnd
