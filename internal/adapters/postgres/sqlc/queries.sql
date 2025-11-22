-- name: ListProducts :many
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;

-- name: CreateOrder :one
INSERT INTO orders (
    customer_id
) VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity, price_in_rubles)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateProductStockQuantity :exec
UPDATE products
SET quantity = quantity - $1
WHERE id = $2;