-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1
LIMIT 1;

-- name: GetAllProducts :many
SELECT * FROM products
ORDER BY created_at DESC;

-- name: CreateProduct :one
INSERT INTO products (title, description, price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProduct :exec
UPDATE products SET
    title = $2,
    description = $3,
    price = $4
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;
