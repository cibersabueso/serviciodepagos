// internal/bank/simulator.go
package bank

import "errors"

func SimulateBankTransaction(amount float64, cardInfo string) (bool, error) {

	if cardInfo == "" {
		// Simula un error si la información de la tarjeta está vacía.
		return false, errors.New("información de la tarjeta inválida")
	}

	// Simulación de una validación del monto.
	if amount <= 0 {
		// Simula un error si el monto es menor o igual a cero.
		return false, errors.New("monto de la transacción inválido")
	}

	// Si llegamos aquí, asumimos que la transacción es exitosa.
	return true, nil
}
