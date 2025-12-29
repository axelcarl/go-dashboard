package query

import "go-dashboard/internal/application/common"

type GetPaymentByIDQuery struct {
	ID int
}

type GetPaymentByIdQueryResult struct {
	Result *common.PaymentResult
}

type GetPaymentsQueryResult struct {
	Result []*common.PaymentResult
}
