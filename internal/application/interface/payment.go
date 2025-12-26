package interfaces

import "go-dashboard/internal/application/query"

type PaymentService interface {
	FindPaymentByID(query *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error)
}
