package models

import ()

type PointsOfSale struct {
	Paging Paging `json:"paging"`
	Results Result `json:"results"`
	Status string `json:"status"`
	Uuid string `json:"uuid"`
	UserId int `json:"user_id"`
	Name string `json:"name"`
	FixedAmount bool `json:"fixed_amount"`
	Category int `json:"category"`
	StoreId int `json:"store_id"`
	ExternalStoreID string `json:"external_store_id"`
	ExternalId string `json:"external_id"`
}

type Result struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Category int `json:"category"`
	StoreId int `json:"store_id"`
	ExternalId int `json:"external_id"`
	Id int `json:"id"`
	QR QR `json:"qr"`
	DateCreated string `json:"date_created"`
	DateLastUpdated string `json:"date_last_updated"`
}

type Paging struct {
	Total int `json:"total"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type QR struct {
	Image string `json:"image"`
	TemplateDocument string `json:"template_document"`
	TemplateImage string `json:"template_image"`
}