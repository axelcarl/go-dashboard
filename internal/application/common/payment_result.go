package common

import "time"

type PaymentResult struct {
	ID        int
	Sender    string
	Recipient string
	Amount    float64
	UpdatedAt time.Time
	CreatedAt time.Time
}
