// internal/handlers/paymentHandlers.go
package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/enrique/goone/internal/payments"
)

func ProcessPaymentHandler(paymentService *payments.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        var p payments.Payment
        err := json.NewDecoder(r.Body).Decode(&p)
        if err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        err = paymentService.ProcessPayment(&p)
        if err != nil {
            // Manejar el error adecuadamente
            log.Printf("Error processing payment: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"status": "success"})
    }
}

func GetPaymentDetailsHandler(paymentService *payments.PaymentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        paymentID, err := strconv.ParseInt(vars["id"], 10, 64)
        if err != nil {
            http.Error(w, "ID de pago inválido", http.StatusBadRequest)
            return
        }

        payment, err := paymentService.GetPaymentByID(paymentID)
        if err != nil {
            http.Error(w, "Pago no encontrado", http.StatusNotFound)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(payment); err != nil {
            http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
        }
    }
}