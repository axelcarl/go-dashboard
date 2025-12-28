package mapper

import (
	"go-dashboard/internal/api/response"
	"go-dashboard/internal/application/common"
)

func ToPaymentResponse(payment *common.PaymentResult) *response.PaymentReponse {
	return &response.PaymentReponse{
		ID:        payment.ID,
		Sender:    payment.Sender,
		Recipient: payment.Recipient,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}
}
