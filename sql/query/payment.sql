-- name: GetPaymentByID :one
SELECT * FROM payments
WHERE id = $1 LIMIT 1;
