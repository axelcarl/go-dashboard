package query

import "go-dashboard/internal/application/common"

type GetPaymentsQueryResult struct {
	Result []*common.PaymentResult
}
