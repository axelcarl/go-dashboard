package entity

import "time"

type Payment struct {
	ID        int32
	Sender    string
	Recipient string
	Amount    float64
	UpdatedAt time.Time
	CreatedAt time.Time
}
