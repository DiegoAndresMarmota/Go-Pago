package models

import ()

type DiscountCampaignModel struct {
	ID 				string 	`json:"id"`
	Name 			string 	`json:"name"`
	PercentOff 		float32 `json:"percent_off"`
	AmountOff 		uint16 	`json:"amount_off"`
	CouponAmount 	uint16 	`json:"coupon_amount"`
	CurrencyID 		string 	`json:"currency_id"`
}