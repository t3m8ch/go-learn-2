// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	decimal "github.com/shopspring/decimal"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (title, description, price)
VALUES ($1, $2, $3)
RETURNING id, created_at, title, description, price
`

type CreateProductParams struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.Title, arg.Description, arg.Price)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Title,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getAllProducts = `-- name: GetAllProducts :many
SELECT id, created_at, title, description, price FROM products
ORDER BY created_at DESC
`

func (q *Queries) GetAllProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Title,
			&i.Description,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
SELECT id, created_at, title, description, price FROM products
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetProductById(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProductById, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Title,
		&i.Description,
		&i.Price,
	)
	return i, err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products SET
    title = $2,
    description = $3,
    price = $4
WHERE id = $1
`

type UpdateProductParams struct {
	ID          int64           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Price,
	)
	return err
}