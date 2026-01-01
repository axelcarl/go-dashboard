package interfaces

import (
	"go-dashboard/internal/application/mutation"
	"go-dashboard/internal/application/query"
)

type PaymentService interface {
	FindPaymentByID(query *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error)
	List() (*query.GetPaymentsQueryResult, error)
	Create(mutation *mutation.CreatePaymentMutation) (*mutation.CreatePaymentMutationResult, error)
}
