package pkg

import "github.com/DiegoAndresMarmota/Go-Pago/api/models"

// Payment proporciona acceso a la API para crear pagos.
type Payment struct {
	Schema	models.PayerModel
}

// Create crea un pago
func (p *Payment) Create() {
	
}

// Update actualiza un pago
func (p *Payment) Update() {
	
}

// Get obtiene un pago por ID
func (p *Payment) Get() {
	
}

// Search busca pagos en la base de datos
func (p *Payment) Search() {
	
}

// Cancel cancela la orden de pago
func (p *Payment) Cancel(id string, config interface{}, callback interface{}) {
	DataCancel := map[string]interface{}{
		"id":     id,
		"status": "cancelled",
	}

	p.Update(DataCancel, config, callback)
}

// Capture captura un pago
func (p *Payment) Capture(id string, config interface{}, callback interface{}) {
	DataCapture := map[string]interface{}{
		"id": id,
	}

	p.Capture(DataCapture, config, callback)
}


// Refund reembolsa un pago
func (p *Payment) Refund(id string, config interface{}, callback interface{}) {
	DataRefund := map[string]interface{}{
		"payment_id": id,
	}

	p.Refund(DataRefund, config, callback)
}
