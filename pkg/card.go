package pkg

import (
	"fmt"
	"net/http"
	"github.com/DiegoAndresMarmota/Go-Pago/api/models"
)

const pagoURL = "https://api.gopago.com/v1/"

type Card struct {
	Schema models.Card
}

// All obtiene todas las tarjetas de un usuario guardadas del usuario titular.
func (c *Card) All(userID string) ([]byte, error) {
	if c, err := http.NewRequest("GET", pagoURL, nil)
	err != nil {
		fmt.Errorf("GetAllUsers not found", err)
	} 
	return c, nil
}

// Get obtiene una tarjeta de un usuario guardada del usuario titular, según su id.
func (c *Card) Get(userID string, cardID string) {

}

// FindID busca la tarjeta guardada del usuario titular, según su id.
func (c *Card) FindID(userID string, cardID string) {

}

// Create crea una nueva tarjeta para un usuario titular.
func (c *Card) Create(userID string) {
}

// Save crea una nueva tarjeta para un usuario titular.
func (c *Card) Save(userID string) {

}

// Update actualiza la información de una tarjeta de un usuario titular.
func (c *Card) Update(userID string, cardID string) {

}

// Delete elimina una tarjeta de un usuario titular.
func (c *Card) Delete(userID string, cardID string) {

}
