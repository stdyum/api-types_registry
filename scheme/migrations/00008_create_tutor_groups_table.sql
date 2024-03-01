-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS tutor_groups
(
    study_place_id uuid not null,
    teacher_id     uuid not null default uuid_generate_v4(),
    group_id       uuid not null,

    CONSTRAINT fk_teacher FOREIGN KEY (teacher_id)
        REFERENCES teachers (id),
    CONSTRAINT fk_group FOREIGN KEY (group_id)
        REFERENCES groups (id),

    UNIQUE (teacher_id, group_id)
);

CALL register_updated_at_created_at_columns('tutor_groups');

-- +goose StatementEnd
