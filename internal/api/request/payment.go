package request

import (
	"errors"
	"go-dashboard/internal/application/mutation"
)

var (
	ErrSenderEmpty    = errors.New("sender is empty")
	ErrRecipientEmpty = errors.New("recipient is empty")
	ErrInvalidAmount  = errors.New("amount must be greater than 0")
)

type PaymentRequest struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

func (body *PaymentRequest) ToCreatePaymentMutation() (*mutation.CreatePaymentMutation, error) {
	if body.Sender == "" {
		return nil, ErrSenderEmpty
	}

	if body.Recipient == "" {
		return nil, ErrRecipientEmpty
	}

	if body.Amount <= 0 {
		return nil, ErrInvalidAmount
	}

	mut := &mutation.CreatePaymentMutation{
		Sender:    body.Sender,
		Recipient: body.Recipient,
		Amount:    body.Amount,
	}

	return mut, nil
}
