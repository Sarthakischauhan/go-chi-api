-- name: GetProducts :many
SELECT * FROM products;

-- name: InsertProducts :one
INSERT INTO products (id, name, price)
VALUES ($1, $2, $3) RETURNING *;