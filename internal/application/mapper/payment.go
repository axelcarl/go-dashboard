package mapper

import (
	"go-dashboard/internal/application/common"
	"go-dashboard/internal/domain/entity"
)

func NewPaymentResultFromEntity(payment *entity.Payment) *common.PaymentResult {
	if payment == nil {
		return nil
	}

	return &common.PaymentResult{
		ID:        int(payment.ID),
		Sender:    payment.Sender,
		Recipient: payment.Recipient,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}
