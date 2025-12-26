package response

import "time"

type PaymentReponse struct {
	ID        int       `json:"id"`
	Sender    string    `json:"sender"`
	Recipient string    `json:"recipient"`
	Amount    float64   `json:"amount"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
