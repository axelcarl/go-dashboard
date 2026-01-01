package repository

import entities "go-dashboard/internal/domain/entity"

type PaymentRepository interface {
	FindByID(id int) (*entities.Payment, error)
	List() ([]*entities.Payment, error)
	Create(payment *entities.ValidatedPayment) (*entities.Payment, error)
}
