// internal/payments/service.go
package payments

type Service interface {
	CreatePayment(p *Payment) error
	GetPaymentByID(id int64) (*Payment, error)
	UpdatePayment(p *Payment) error
	ProcessPayment(p *Payment) error
}
