// internal/handlers/customer_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/enrique/goone/internal/customers"
)

func RegisterCustomerHandler(customerService *customers.CustomerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newCustomer customers.Customer
		if err := json.NewDecoder(r.Body).Decode(&newCustomer); err != nil {
			http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
			return
		}

		err := customerService.CreateCustomer(&newCustomer)
		if err != nil {
			http.Error(w, "Error al registrar el cliente", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCustomer)
	}
}