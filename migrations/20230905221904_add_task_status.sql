-- +goose Up
-- +goose StatementBegin
CREATE TYPE task_status AS ENUM ('todo', 'in_progress', 'done');
ALTER TABLE tasks ADD COLUMN status task_status;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS task_status;
ALTER TABLE tasks DROP COLUMN status;
-- +goose StatementEnd
