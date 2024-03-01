// internal/payments/payment.go
package payments

import "time"

type Payment struct {
	PaymentID  int64     `json:"payment_id"`
	MerchantID int64     `json:"merchant_id"`
	CustomerID int64     `json:"customer_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
