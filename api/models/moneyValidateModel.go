package models

import ()

type MoneyValidate struct {
	ConceptType []ConceptEnum
	Amount int64 `json:"amount"`
	Description string `json:"description"`
	CurrencyId []CurrencyEnum
	PayerEmail string `json:"payer_email"`
	Status []StatusEnum
	CollectorID int64 `json:"collector_id"`
	Id string `json:"id"`
	SiteId []SiteEnum
}

type ConceptEnum struct {
	OnPlataform bool `json:"on_plataform"`
	OffPlataform bool `json:"off_plataform"`
}

type CurrencyEnum struct {
	ARS string `json:"ARS"`
	BRL string `json:"BRL"`
	MXN string `json:"MXN"`
	VEF string `json:"VEF"`
	COP string `json:"COP"`
	CLP string `json:"CLP"`
}

type StatusEnum struct {
	Pending string `json:"pending"`
	Accepted string `json:"accepted"`
	Rejected string `json:"rejected"`
	Cancelled string `json:"cancelled"`
}

type SiteEnum struct {
	MLA string `json:"MLA"`
	MLB string `json:"MLB"`
	MLM string `json:"MLM"`
	MLV string `json:"MLV"`
	MCO string `json:"MCO"`
	MLC string `json:"MLC"`
}