// internal/merchants/merchant.go
package merchants

import "time"

type Merchant struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	MerchantCode string    `json:"merchant_code"`
	CreatedAt    time.Time `json:"created_at"`
}
