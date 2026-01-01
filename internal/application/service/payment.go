package service

import (
	"go-dashboard/internal/application/mapper"
	"go-dashboard/internal/application/mutation"
	"go-dashboard/internal/application/query"
	"go-dashboard/internal/domain/entity"
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
	storedPayment, err := s.paymentRepository.FindByID(q.ID)
	if err != nil {
		return nil, err
	}

	var queryResult query.GetPaymentByIdQueryResult
	queryResult.Result = mapper.NewPaymentResultFromEntity(storedPayment)

	return &queryResult, nil
}

func (s *PaymentService) List() (*query.GetPaymentsQueryResult, error) {
	storedPayments, err := s.paymentRepository.List()
	if err != nil {
		return nil, err
	}

	var queryResult query.GetPaymentsQueryResult
	for _, payment := range storedPayments {
		queryResult.Result = append(queryResult.Result, mapper.NewPaymentResultFromEntity(payment))
	}

	return &queryResult, nil
}

func (s *PaymentService) Create(m *mutation.CreatePaymentMutation) (*mutation.CreatePaymentMutationResult, error) {
	payment := entity.NewPayment(m.Sender, m.Recipient, m.Amount)
	validatedPayment, err := entity.NewValidatedPayment(payment)
	if err != nil {
		return nil, err
	}

	newPayment, err := s.paymentRepository.Create(validatedPayment)
	if err != nil {
		return nil, err
	}

	var mutationResult mutation.CreatePaymentMutationResult
	mutationResult.Result = mapper.NewPaymentResultFromEntity(newPayment)

	return &mutationResult, err
}
