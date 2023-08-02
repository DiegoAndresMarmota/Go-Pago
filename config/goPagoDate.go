package config

import "time"

type GoPagoDate struct {
	DatePurchase time.Time
}


func NewGoPagoDate(args ...interface{}) GoPagoDate {
	var date time.Time
	if len(args) >=  1 {
		date = time.Date(date.Year(), date.Month(), date.Day(), args[0].(int), args[1].(int), 0, 0, time.Local)
	} else if len(args) >=  2 {
		date = time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, time.Local)
	} else {
		date = time.Now()
	}
	return GoPagoDate{
		DatePurchase: date,
	}
}

func (gpd GoPagoDate) ToString(utc uint8) string {
	if utc != 0 {
		utcOff := time.FixedZone("UTC", int(utc))
		return gpd.DatePurchase.In(utcOff).Format("2023-01-02 13:13:13")
	}
	return gpd.DatePurchase.Format("2023-01-02 13:13:13")
}