package models

import ()

type StatusPurchaseModel struct {
	PropPurchase []PropPurchase
}

type PropPurchase struct {
	PayerEmail string `json:"payer_email"`
	BackUrl    string `json:"back_url"`
	StatusPurchase StatusPurchase
	AutoRecurring []AutoRecurring
	Reason string `json:"reason"`
	ExternalReference string `json:"external_reference"`
}

type StatusPurchase interface {
	Pending()
	Authorized()
	Paused()
	Cancelled()
}

type AutoRecurring struct {
	Frequency uint16 `json:"frequency"`
	FrequencyType FreqYear
	TransactionAmount uint16 `json:"transaction_amount"`
	CurrencyID string `json:"currency_id"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}

type FreqYear interface {
	Days()
	Months()
}

