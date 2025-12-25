package repository

import (
	"database/sql"

	_ "github.com/joho/godotenv/autoload"

	"go-dashboard/internal/generated/sqlc"
)

type PaymentRepository struct {
	Query *sqlc.Queries
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		Query: sqlc.New(db),
	}
}
