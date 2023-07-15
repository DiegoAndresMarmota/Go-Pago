package models

import ()

type PayMerchantModel struct {
	PreferenceId string `json:"preference_id"`
	AplicationId string `json:"application_id"`
	SiteId string `json:"site_id"`
	Payer PayerMerchant
	Collector CollectorMerchant
	SponsorId int `json:"sponsor_id"`
	Cancelled bool `json:"cancelled"`
	Shipments ShipmentMerchant
	NotificationUrl string `json:"notification_url"`
	AdditionalInfo string `json:"additional_info"`
	ExternalReference string `json:"external_reference"`
	MarketPlace string `json:"marketplace"`
}

type PayerMerchant struct {
	Id string `json:"id"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
}

type CollectorMerchant struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
}

type ShipmentMerchant struct {
	Id int `json:"id"`
	ShipmentType string `json:"shipment_type"`
	ShippingMode string `json:"shipment_mode"`
	PickingType string `json:"picking_type"`
	Status string `json:"status"`
	SubStatus string `json:"sub_status"`
	Items ItemsShipment
	DateCreated string `json:"date_created"`
	LastModified string `json:"last_modified"`
	DateFirstPrinted string `json:"date_first_printed"`
	ServiceId string `json:"service_id"`
	SenderId int `json:"sender_id"`
	ReceiverId int `json:"receiver_id"`
	ReceiverAddress ReceiverAddress
}

type ItemsShipment struct {
	Id string `json:"id"`
	CategoryId string `json:"category_id"`
	CurrencyId string `json:"currency_id"`
	Description string `json:"description"`
	PictureUrl string `json:"picture_url"`
	Quantity int `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Title string `json:"title"`
}

type ReceiverAddress struct {
	ZipCode string `json:"zip_code"`
	StreetName string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Floor string `json:"floor"`
	Apartment string `json:"apartment"`
}