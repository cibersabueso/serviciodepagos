// internal/refunds/service.go
package refunds

type Service interface {
	CreateRefund(r *Refund) error
	GetRefundByID(id int64) (*Refund, error)
	UpdateRefund(r *Refund) error
	ProcessRefund(r *Refund) error
}
