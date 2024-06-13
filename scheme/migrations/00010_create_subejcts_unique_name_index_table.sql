-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX idx_subjects_name_unique ON subjects (study_place_id, name)

-- +goose StatementEnd
