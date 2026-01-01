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

func (repo *SqlcPaymentRepository) List() ([]*entity.Payment, error) {
	ctx := context.Background()

	rows, err := repo.queries.GetPayments(ctx)
	if err != nil {
		return nil, err
	}

	payments := make([]*entity.Payment, len(rows))
	for i, row := range rows {
		payments[i] = fromSqlcPaymentRow(&row)
	}

	return payments, nil
}

func (repo *SqlcPaymentRepository) Create(validatedPayment *entity.ValidatedPayment) (*entity.Payment, error) {
	ctx := context.Background()
	strAmount := strconv.FormatFloat(validatedPayment.Amount, 'f', 3, 64)

	params := sqlc.CreatePaymentParams{Sender: validatedPayment.Sender, Recipient: validatedPayment.Recipient, Amount: strAmount}
	payment, err := repo.queries.CreatePayment(ctx, params)
	if err != nil {
		return nil, err
	}

	return fromSqlcPaymentRow(&payment), nil
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
