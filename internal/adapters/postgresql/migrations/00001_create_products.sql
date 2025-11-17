-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGERS NOT NULL,
    created_at TIMESTAMPZ NOT NULL DEFAULT now()
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products
-- +goose StatementEnd
