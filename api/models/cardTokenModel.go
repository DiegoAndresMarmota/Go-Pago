package models

import ()

type CardToken struct {
	ID string `json:"id"`
	CustomerID string `json:"customer_id"`
	CardID string `json:"card_id"`
	SecurityCode string `json:"security_code"`
	Status string `json:"status"`
	DateCreated string `json:"date_created"`
	DateLastUpdated string `json:"date_last_updated"`
	DateDue string `json:"date_due"`
	LuhnValidation bool `json:"luhn_validation"`
	LiveMode bool `json:"live_mode"`
	RequireESC bool `json:"require_esc"`
}