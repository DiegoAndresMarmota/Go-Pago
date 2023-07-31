package config

import "fmt"

type GoPagoResponse struct {
	ID 		string
	Body 	[]byte
	Status 	int
	Topic   string
}

type GoPagoPrintResponse struct {
	IDResponse 		string
	BodyResponse 	string
	StatusResponse 	string
	TopicResponse 	string
}

type GoPagoStructResponse struct {
	GoPagoResponse
	GoPagoPrintResponse
}

func NewGoPagoStructResponse(id string, body []byte, status int, topic string) *GoPagoStructResponse {
	return &GoPagoStructResponse{
		GoPagoResponse: GoPagoResponse || GoPagoPrintResponse{
			ID: id,     				//"12345"
			IDResponse: fmt.Printf("ResponseID: %s\n", id),
			Body: body,					//{"key": "value"}
			BodyResponse: fmt.Printf("ResponseBody: %s\n", string(body)), 				
			Status: status,				//"canceled"
			StatusResponse: fmt.Printf("ResponseStatus: %s\n", status),
			Topic: topic,				//"payment"			
			TopicResponse: fmt.Printf("ResponseTopic: %s\n", topic),
		}
	}
}
