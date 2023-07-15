package models

import ()

type Card struct {
	Id                string `json:"id"`
	Token             string `json:"token"`
	UserId 			  string `json:"customer_id"`
	ExpirationMonth   int    `json:"expiration_month"`
	ExpirationYear    int    `json:"expiration_year"`
	FirstSixDigits    string `json:"first_six_digits"`
	LastFourDigits    string `json:"last_four_digits"`
	PaymentProperties PaymentProperties
	SecurityCode      SecurityCode
	Issuer            Issuer
	Cardholder        Cardholder
	DateCreated       string `json:"date_created"`
	DateLastUpdated   string `json:"date_last_updated"`
	PaymentMethodId   string `json:"payment_method_id"`
}

type PaymentProperties struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	PaymentTypeId   string `json:"payment_type_id"`
	Thumbnail       string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

type SecurityCode struct {
	Length       int    `json:"length"`
	CardLocation string `json:"card_location"`
}

type Issuer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Cardholder struct {
	Name           string `json:"name"`
	Identification IdentificationHolder
}

type IdentificationHolder struct {
	Number  string `json:"number"`
	SubType string `json:"sub_type"`
	Type    string `json:"type"`
}
