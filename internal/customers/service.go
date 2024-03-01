// internal/customers/service.go
package customers

type Service interface {
	CreateCustomer(c *Customer) error
	GetCustomerByID(id int64) (*Customer, error)
	UpdateCustomer(c *Customer) error
	DeleteCustomer(id int64) error
}
