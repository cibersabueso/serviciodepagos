// cmd/api/main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/enrique/goone/internal/customers"
	"github.com/enrique/goone/internal/database"
	"github.com/enrique/goone/internal/handlers"
	"github.com/enrique/goone/internal/merchants"
	"github.com/enrique/goone/internal/payments"
	"github.com/enrique/goone/internal/refunds"
)

type PaymentRequest struct {
	MerchantID int64   `json:"merchant_id,string"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status,omitempty"`
	CreatedAt  string  `json:"created_at"`
}

func main() {
	dbPool, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	paymentService := payments.NewPaymentService(dbPool)
	merchantService := merchants.NewMerchantService(dbPool)
	merchantHandler := handlers.NewMerchantHandler(merchantService)
	customerService := customers.NewCustomerService(dbPool)
	refundService := refunds.NewRefundService(dbPool)

	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/process-payment", handlers.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req PaymentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Error al decodificar la solicitud de pago", http.StatusBadRequest)
			return
		}

		customerID, err := strconv.ParseInt(req.CustomerID, 10, 64)
		if err != nil {
			http.Error(w, "Error al convertir CustomerID a int64", http.StatusBadRequest)
			return
		}

		payment := payments.Payment{
			MerchantID: req.MerchantID,
			CustomerID: customerID,
			Amount:     req.Amount,
			Status:     req.Status,
		}

		err = paymentService.ProcessPayment(&payment)
		if err != nil {
			http.Error(w, "Error al procesar el pago", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pago procesado exitosamente"))
	}))).Methods("POST")

	r.Handle("/merchants", merchantHandler).Methods("POST", "PUT")
	r.Handle("/merchants/{id:[0-9]+}", merchantHandler).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/customers", handlers.RegisterCustomerHandler(customerService)).Methods("POST")
	r.HandleFunc("/payments/{id:[0-9]+}", handlers.GetPaymentDetailsHandler(paymentService)).Methods("GET")

	r.HandleFunc("/refunds", handlers.ProcessRefundHandler(refundService)).Methods("POST")

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
