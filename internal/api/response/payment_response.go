package response

import "time"

type PaymentReponse struct {
	ID        int
	Sender    string
	Recipient string
	Amount    float64
	UpdatedAt time.Time
	CreatedAt time.Time
}
