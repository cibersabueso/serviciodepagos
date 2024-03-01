// internal/payments/paymentservice.go
package payments

import (
	"context"
	"fmt"
	"log"

	"github.com/enrique/goone/internal/bank" // Corregido: Ruta correcta del paquete bank
	"github.com/jackc/pgx/v4/pgxpool"
)

type PaymentService struct {
	db *pgxpool.Pool
}

func NewPaymentService(db *pgxpool.Pool) *PaymentService {
	return &PaymentService{db: db}
}

// VerifyMerchantExists verifica si un merchant_id existe en la base de datos.
func (s *PaymentService) VerifyMerchantExists(merchantID int64) (bool, error) {
	const query = `SELECT COUNT(*) FROM merchants WHERE merchant_id = $1`
	var count int
	err := s.db.QueryRow(context.Background(), query, merchantID).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar el merchant_id: %v\n", err)
		return false, err
	}
	return count > 0, nil
}

// Añadido: Método para verificar la existencia de un customer_id
func (s *PaymentService) VerifyCustomerExists(customerID int64) (bool, error) {
	const query = `SELECT COUNT(*) FROM customers WHERE customer_id = $1`
	var count int
	err := s.db.QueryRow(context.Background(), query, customerID).Scan(&count)
	if err != nil {
		log.Printf("Error al verificar el customer_id: %v\n", err)
		return false, err
	}
	return count > 0, nil
}

func (s *PaymentService) ProcessPayment(p *Payment) error {
	log.Println("Iniciando el procesamiento del pago")

	// Verificar si el merchant_id existe
	exists, err := s.VerifyMerchantExists(p.MerchantID)
	if err != nil {
		log.Printf("Error al verificar el merchant_id: %v\n", err)
		return err
	}
	if !exists {
		errMsg := fmt.Sprintf("El merchant_id %d no existe", p.MerchantID)
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	// Añadido: Verificar si el customer_id existe
	customerExists, err := s.VerifyCustomerExists(p.CustomerID)
	if err != nil {
		log.Printf("Error al verificar el customer_id %d: %v\n", p.CustomerID, err)
		return err
	}
	if !customerExists {
		errMsg := fmt.Sprintf("El customer_id %d no existe", p.CustomerID)
		log.Println(errMsg)
		return fmt.Errorf(errMsg)
	}

	// Simulación de la interacción bancaria
	success, err := bank.SimulateBankTransaction(p.Amount, "información de tarjeta")
	if err != nil {
		log.Printf("Error en la simulación bancaria: %v\n", err)
		return err
	}
	if !success {
		p.Status = "failed"
	} else {
		p.Status = "success"
	}

	// Insertar el pago en la base de datos
	const insertQuery = `
        INSERT INTO payments (merchant_id, customer_id, amount, status, created_at)
        VALUES ($1, $2, $3, $4, NOW())
        RETURNING payment_id`

	log.Printf("Insertando pago: %+v\n", p)
	err = s.db.QueryRow(context.Background(), insertQuery, p.MerchantID, p.CustomerID, p.Amount, p.Status).Scan(&p.PaymentID)
	if err != nil {
		log.Printf("Error al insertar el pago: %v\n", err)
		return err
	}

	log.Println("Pago procesado exitosamente")
	return nil
}

func (s *PaymentService) GetPaymentByID(id int64) (*Payment, error) {
	var p Payment
	const query = `SELECT payment_id, merchant_id, customer_id, amount, status, created_at FROM payments WHERE payment_id = $1`
	err := s.db.QueryRow(context.Background(), query, id).Scan(&p.PaymentID, &p.MerchantID, &p.CustomerID, &p.Amount, &p.Status, &p.CreatedAt)
	if err != nil {
		log.Printf("Error al recuperar el pago con payment_id %d: %v\n", id, err)
		return nil, err
	}
	return &p, nil
}
