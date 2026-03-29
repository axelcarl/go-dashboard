package mutation

import "go-dashboard/internal/application/common"

type UpdatePaymentMutation struct {
	ID        int
	Sender    string
	Recipient string
	Amount    float64
}

type UpdatePaymentMutationResult struct {
	Result *common.PaymentResult
}
