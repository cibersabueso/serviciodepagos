// internal/merchants/service.go
package merchants

type Service interface {
	CreateMerchant(m *Merchant) error
	GetMerchantByID(id int64) (*Merchant, error)

	UpdateMerchant(id int64, m *Merchant) error
	DeleteMerchant(id int64) error
}
