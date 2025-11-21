-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE TABLE IF NOT EXISTS order_products(
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL, 
    product_id BIGINT NOT NULL, 
    quantity INTEGER NOT NULL, 
    price INTEGER NOT NULL, 
    CONSTRAINT fk_order FOREIGN KEY (order_id) references orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS order_products;
-- +goose StatementEnd
