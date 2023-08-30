-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS images (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    image_data text NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;
-- +goose StatementEnd
