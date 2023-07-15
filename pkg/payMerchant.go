package pkg

import (
	"github.com/DiegoAndresMarmota/Go-Pago/api/models"
)

// PayMerchant representa la orden de pago al comerciante.
type PayMerchant struct {
	Schema models.PayMerchantModel
}

// Create crea una orden de pago al comerciante.
func (pm *PayMerchant) Create() (_, error) {
	return
}

// Save guarda una orden de pago al comerciante.
func (pm *PayMerchant) Save() (_, error) {
	return pm.Create()
}

// Update actualiza una orden de pago al comerciante.
func (pm *PayMerchant) Update(id string) (_, error) {
	return
}

// Get obtiene una orden de pago al comerciante, según su ID.
func (pm *PayMerchant) Get(id string) (_, error) {
	return
}
