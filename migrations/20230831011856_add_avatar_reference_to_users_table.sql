-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD CONSTRAINT fk_avatar FOREIGN KEY (avatar_id) REFERENCES images(id) ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP CONSTRAINT fk_avatar;
-- +goose StatementEnd
