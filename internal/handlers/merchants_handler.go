// internal/handlers/merchants_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/enrique/goone/internal/merchants"
	"github.com/gorilla/mux"
)

type MerchantHandler struct {
	Service merchants.Service
}

func NewMerchantHandler(service merchants.Service) *MerchantHandler {
	return &MerchantHandler{
		Service: service,
	}
}

func (h *MerchantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createMerchant(w, r)
	case http.MethodGet:
		h.getMerchantByID(w, r)
	case http.MethodPut:
		h.updateMerchant(w, r)
	case http.MethodDelete:
		h.deleteMerchant(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *MerchantHandler) createMerchant(w http.ResponseWriter, r *http.Request) {
	var m merchants.Merchant
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.CreateMerchant(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

func (h *MerchantHandler) getMerchantByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	m, err := h.Service.GetMerchantByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(m)
}

func (h *MerchantHandler) updateMerchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var m merchants.Merchant
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateMerchant(id, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func (h *MerchantHandler) deleteMerchant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteMerchant(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Comerciante eliminado exitosamente"))
}
