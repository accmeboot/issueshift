-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email citext NOT NULL,
    password_hash bytea NOT NULL,
    name text NOT NULL,
    created_at timestamp DEFAULT now(),
    avatar_url text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
