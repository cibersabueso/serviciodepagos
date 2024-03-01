package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/enrique/goone/internal/auth" // Ajusta el import
)

type User struct {
	Username string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los datos del usuario", http.StatusBadRequest)
		return
	}

	tokenString, err := auth.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Error al generar el token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
