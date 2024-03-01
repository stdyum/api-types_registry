-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS student_groups
(

    study_place_id uuid not null,
    student_id     uuid not null default uuid_generate_v4(),
    group_id       uuid not null,

    CONSTRAINT fk_student FOREIGN KEY (student_id)
        REFERENCES students (id),
    CONSTRAINT fk_group FOREIGN KEY (group_id)
        REFERENCES groups (id),

    UNIQUE (student_id, group_id)
);

CALL register_updated_at_created_at_columns('student_groups');

-- +goose StatementEnd
