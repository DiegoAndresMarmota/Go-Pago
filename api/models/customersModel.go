package models

import ()

type CustomersModel struct {
	Email    		string `json:"email"`
	FirstName 		string `json:"firstName"`
	LastName  		string `json:"lastName"`
	Phone 			[]PhoneCustomer
	Identification 	[]IdentificationCustomer
	AddressCustomer	string `json:"default_address"`
	Address 		[]Address
	DateRegistered 	string `json:"date_registered"`
	Description 	string `json:"description"`
	Metadata 		string `json:"metadata"`
	DefaultCard 	string `json:"default_card"`
	Cards 			[]Cards
	Addresses		[]Addresses
}

type PhoneCustomer struct {
	AreaCode string `json:"areaCode"`
	Number   string `json:"number"`
}

type IdentificationCustomer struct {
	Type 	string `json:"type"`
	Number 	string `json:"number"`
}

type Address struct {
	ZipCode 		string `json:"zip_code"`
	StreetName 		string `json:"street_name"`
	StreetNumber 	uint16 `json:"street_number"`
}

type Cards struct {
	CustomerId 					string `json:"customer_id"`
	ExpirationMonth 			string `json:"expiration_month"`
	ExpirationYear 				string `json:"expiration_year"`
	FirstSixDigitals 			string `json:"first_six_digitals"`
	LastFourDigits 				string `json:"last_four_digits"`
	PaymentMethodCustomer 		[]PaymentMethodCustomer
	SecurityCode				[]SecurityCode
	Issuer 						[]Issuer
	Cardholder 					[]Cardholder
}

type PaymentMethodCustomer struct {
	ID 				uint16 `json:"id"`
	Name 			string `json:"name"`
	PaymentTypeID 	string `json:"payment_type_id"`
	Thumbnail 		string `json:"thumbnail"`
	SecureThumbnail string `json:"secure_thumbnail"`
}

type Addresses struct {
	Phone 			string 	`json:"phone"`
	Name 			string 	`json:"name"`
	Floor 			uint8 	`json:"floor"`
	Apartment 		string `json:"apartment"`
	StreetName 		string `json:"street_name"`
	StreetNumber 	uint16 `json:"street_number"`
	ZipCode 		uint16 `json:"zip_code"`
	City 			[]City
	State 			[]State
	Country 		[]Country
	Neighborhood 	[]Neighborhood
	Municipality 	[]Municipality
	Comments 		string `json:"comments"`			
}

type City struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
}

type State struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
}

type Country struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
}

type Neighborhood struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
}

type Municipality struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
}