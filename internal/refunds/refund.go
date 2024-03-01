// internal/refunds/refund.go
package refunds

import "time"

type Refund struct {
	ID        int64     `json:"id"`
	PaymentID int64     `json:"payment_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
