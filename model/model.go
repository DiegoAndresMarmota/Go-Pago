package model
import ()

type MessageResponse struct {
	Error   string `json:"error"`   
	Message string `json:"message"` 
	Status  int    `json:"status"`
}