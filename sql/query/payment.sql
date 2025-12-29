-- name: GetPaymentByID :one
SELECT * FROM payments
WHERE id = $1 LIMIT 1;

-- name: GetPayments :many
SELECT * FROM payments
ORDER BY created_at DESC;
