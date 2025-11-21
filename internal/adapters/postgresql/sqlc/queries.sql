-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;

-- name: InsertProducts :one
INSERT INTO products (id, name, price)
VALUES ($1, $2, $3) RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (customer_id, created_at)
values ($1, $2) RETURNING *;

-- name: CreateOrderProducts :one
INSERT INTO order_products (order_id, product_id, quantity, price)
values ($1, $2, $3, $4) RETURNING *;