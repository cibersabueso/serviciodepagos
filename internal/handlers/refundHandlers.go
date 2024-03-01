// internal/handlers/refundHandlers.go
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/enrique/goone/internal/refunds"
)

func ProcessRefundHandler(refundService *refunds.RefundService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			PaymentID int64   `json:"payment_id"`
			Amount    float64 `json:"amount"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Error al decodificar la solicitud: %v", err)
			http.Error(w, "Error al decodificar la solicitud", http.StatusBadRequest)
			return
		}

		refund := refunds.Refund{
			PaymentID: req.PaymentID,
			Amount:    req.Amount,
			Status:    "processing",
		}

		err := refundService.ProcessRefund(&refund)
		if err != nil {
			log.Printf("Error al procesar el reembolso para el PaymentID %d: %v", req.PaymentID, err)
			http.Error(w, "Error al procesar el reembolso", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(refund); err != nil {
			log.Printf("Error al codificar la respuesta del reembolso: %v", err)

		}
	}
}
