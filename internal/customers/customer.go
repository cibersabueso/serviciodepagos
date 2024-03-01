// internal/customers/customer.go
package customers

import "time"

type Customer struct {
	CustomerID int64     `json:"customer_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}
