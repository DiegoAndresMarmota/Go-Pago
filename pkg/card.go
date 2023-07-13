package pkg 


type Card struct {
	Schema Card
}

// All obtiene todas las tarjetas de un usuario guardadas del usuario titular.
func (c *Card) All(userID string) {

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
