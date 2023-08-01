package config

import ()

type GoPagoError struct {
	name      		string
	message   		string
	status    		uint8
	Options   		ErrorOptions
	AdminRequest	AdminRequest
}

type ErrorOptions struct {
	Error uint8
	Message string
}

type AdminRequest interface {
	AccessTokenResponse() *ErrorOptions
}

func NewMessageGoPagoError(name, message string, status uint8) *GoPagoError {
	return &GoPagoError{
		name: 		name,
		message: 	message,
		status: 	status,
	}
}


func (e *GoPagoError) ValueError() string {
	return e.message
}
