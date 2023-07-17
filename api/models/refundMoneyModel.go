package models

import ()

type RefundMoney struct {
	Id string `json:"id"`
	PaymentId string `json:"payment_id"`
	Amount int64 `json:"amount"`
	Source SourceProperties
	DateCreated string `json:"date_created"`
	UniqueSequenceNumber string `json:"unique_sequence_number"`
}

type SourceProperties struct {
	SourceId string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}