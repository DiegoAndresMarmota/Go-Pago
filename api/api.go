package api

import "fmt"


const BASEURL = "https://api.mercadopago.com"

//
func getenvToken(OwnToken ...string) string {
	if len(getenvToken) == nil {
		return fmt.Println{"Ingrese Token de acceso"}
	} else {
		return os.Getenv("yourAccessTokenMercadoPago")
	}
}

//
func CreatePOS(Data, yourAccessTokenMercadoPago ...string)(*MessageResponse, error){

	params := config.Settings{
		URL: BASEURL + "/pos",
		Method:  "POST",
		Headers: {"Authorization": "Bearer "+ envToken(yourAccessTokenMercadoPago)}, {"Content-Type: application/json"},
		Data: {
			"name": "First POS",
			"fixed_amount": false, 
			"store_id": 1234567, 
			"external_store_id": "SUC001", 
			"external_id": "SUC001POS001", 
			"category": 621102,
			},
	}

	response, err := config.New(Settings)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 400 {
		resp, err := parseError {
			fmt.Println("MISSING_BODY: "HTTP body is required to create pos.)
		}
	}