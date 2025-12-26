package db

import (
	"context"
	"go-dashboard/internal/domain/entity"
	"go-dashboard/internal/domain/repository"
	"go-dashboard/internal/generated/sqlc"
	"strconv"
)

type SqlcPaymentRepository struct {
	queries *sqlc.Queries
}

func NewSqlcPaymentRepository(queries *sqlc.Queries) repository.PaymentRepository {
	return &SqlcPaymentRepository{queries: queries}
}

func (repo *SqlcPaymentRepository) FindByID(id int) (*entity.Payment, error) {
	ctx := context.Background()

	row, err := repo.queries.GetPaymentByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return fromSqlcPaymentRow(&row), nil
}

func fromSqlcPaymentRow(row *sqlc.Payment) *entity.Payment {
	amount, _ := strconv.ParseFloat(row.Amount, 64)
	payment := &entity.Payment{
		Sender:    row.Sender,
		Recipient: row.Recipient,
		Amount:    amount,
		UpdatedAt: row.UpdatedAt,
		CreatedAt: row.CreatedAt,
	}
	payment.ID = row.ID

	return payment
}
