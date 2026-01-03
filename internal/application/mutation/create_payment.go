package mutation

import "go-dashboard/internal/application/common"

type CreatePaymentMutation struct {
	Sender    string
	Recipient string
	Amount    float64
}

type CreatePaymentMutationResult struct {
	Result *common.PaymentResult
}
