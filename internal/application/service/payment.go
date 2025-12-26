package service

import (
	"go-dashboard/internal/application/mapper"
	"go-dashboard/internal/application/query"
	"go-dashboard/internal/domain/repository"
)

type PaymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
	}
}

func (s *PaymentService) FindPaymentByID(q *query.GetPaymentByIDQuery) (*query.GetPaymentByIdQueryResult, error) {
	storedProduct, err := s.paymentRepository.FindByID(q.ID)
	if err != nil {
		return nil, err
	}

	var queryResult query.GetPaymentByIdQueryResult
	queryResult.Result = mapper.NewPaymentResultFromEntity(storedProduct)

	return &queryResult, nil
}
