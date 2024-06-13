-- +goose Up
-- +goose StatementBegin

CREATE UNIQUE INDEX idx_groups_name_unique ON groups (study_place_id, name)

-- +goose StatementEnd
