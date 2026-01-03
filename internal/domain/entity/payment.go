package entity

import (
	"errors"
	"time"
)

type Payment struct {
	ID        int32
	Sender    string
	Recipient string
	Amount    float64
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (p *Payment) validate() error {
	if p.Sender == "" {
		return errors.New("sender must not be empty")
	}

	if p.Recipient == "" {
		return errors.New("Recipient must not be empty")
	}

	if p.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if p.CreatedAt.After(p.UpdatedAt) {
		return errors.New("created_at must be before updated_at")
	}

	return nil
}

func NewPayment(sender string, recipient string, amount float64) *Payment {
	return &Payment{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}
