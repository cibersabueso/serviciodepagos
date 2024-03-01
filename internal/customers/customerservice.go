package customers

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CustomerService struct {
	db *pgxpool.Pool
}

func NewCustomerService(db *pgxpool.Pool) *CustomerService {
	return &CustomerService{db: db}
}

// CreateCustomer implementa la lógica para crear un cliente en la base de datos.
func (s *CustomerService) CreateCustomer(c *Customer) error {
	query := "INSERT INTO customers (name, email, created_at) VALUES ($1, $2, $3) RETURNING customer_id"
	err := s.db.QueryRow(context.Background(), query, c.Name, c.Email, c.CreatedAt).Scan(&c.CustomerID)
	if err != nil {
		log.Printf("error al crear el cliente: %v", err)
		return fmt.Errorf("error al crear el cliente: %w", err)
	}
	return nil
}

// ListCustomers implementa la lógica para listar todos los clientes.
func (s *CustomerService) ListCustomers() ([]Customer, error) {
	var customers []Customer

	query := "SELECT customer_id, name, email, created_at FROM customers"
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar los clientes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.CustomerID, &c.Name, &c.Email, &c.CreatedAt); err != nil { // Cambiado de c.ID a c.CustomerID
			return nil, fmt.Errorf("error al leer el cliente: %w", err)
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error en los resultados: %w", err)
	}

	return customers, nil
}
