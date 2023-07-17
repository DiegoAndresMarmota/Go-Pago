package models

import ()

type Configuration struct {
	Items []ConfigItems
	Tracks []ConfigTracks
	Metadata []ConfigMetadata
	Payer []ConfigPayer
	StatementDescriptor string `json:"statement_descriptor"`
	PaymentMethods []ConfigPaymentMethods
	Shipments []ConfigShipments
	BackUrls []ConfigBackUrls
	Notification map[string]interface{} `json:"notification"`
	Mode []ModeEnum
	AdditionalInfo map[string]interface{} `json:"additional_info"`
	AutoReturn []AutoReturnEnum
	ExternalReference map[string]interface{} `json:"external_reference"`
	Expires bool `json:"expires"`
	ExpirationDateFrom map[string]interface{} `json:"expiration_date_from"`
	ExpirationDateTo map[string]interface{} `json:"expiration_date_to"`
	CollectorId uint64 `json:"collector_id"`
	ClientId uint64 `json:"client_id"`
	Marketplace map[string]interface{} `json:"marketplace"`
	MarketplaceFee uint16 `json:"marketplace_fee"`
	DifferentialPricing int8 `json:"differential_pricing"`
	Taxes []ConfigTaxes
	BinaryMode bool `json:"binary_mode"`
}

type ConfigItems struct {
	Title map[string]interface{} `json:"title"`
	Description map[string]interface{} `json:"description"`
	PictureURL map[string]interface{} `json:"picture_url"`
	CategoryId map[string]interface{} `json:"category_id"`
	Quantity map[uint8]interface{} `json:"quantity"`
	CurrencyId map[string]interface{} `json:"currency_id"`
	UnitPrice map[uint8]interface{} `json:"unit_price"`
}

type ConfigTracks struct {
	Type map[string]interface{} `json:"type"`
	Values []ConfigValues
}

type ConfigValues struct {
	ConversionId string `json:"conversion_id"`
	ConversionLabel string `json:"conversion_label"`
	PixelId string `json:"pixel_id"`
}

type ConfigMetadata struct {}

type ConfigPayer struct {
	Name map[string]interface{} `json:"name"`
	Surname map[string]interface{} `json:"surname"`
	Email map[string]interface{} `json:"email"`
	Phone []PhoneProperties
	Identification []IdentProperties
	Address []AddressProperties

}

type PhoneProperties struct {
	AreaCode map[string]interface{} `json:"area_code"`
	Number map[string]interface{} `json:"number"`
}

type IdentProperties struct {
	Type map[string]interface{} `json:"type"`
	Number map[string]interface{} `json:"number"`
}

type AddressProperties struct {
	ZipCode map[string]interface{} `json:"zip_code"`
	StreetName map[string]interface{} `json:"street_name"`
	StreetNumber uint64 `json:"street_number"`
}

type ConfigPaymentMethods struct {
	ExcludedPaymentMethods map[string]interface{} `json:"excluded_payment_methods"`
	ExcludedPaymentTypes map[string]interface{} `json:"excluded_payment_types"`
	DefaultPaymentMethodId map[string]interface{} `json:"default_payment_method_id"`
	Installments uint64 `json:"installments"`
	DefauktInstallments uint64 `json:"default_installments"`
}

type ConfigShipments struct {
	Mode map[string]interface{} `json:"mode"`
	LocalPickup bool `json:"local_pickup"`
	Dimensions string `json:"dimensions"`
	DefaultShippingMethod uint64 `json:"default_shipping_method"`
	FreeMethods uint16 `json:"free_methods"`
	Cost uint32 `json:"cost"`
	FreeShipping bool `json:"free_shipping"`
	ReceiverAddress []ReceiverAddressProperties
}

type ReceiverAddressProperties struct {
	ZipCode map[string]interface{} `json:"zip_code"`
	StreetName map[string]interface{} `json:"street_name"`
	StreetNumber map[string]interface{} `json:"street_number"`
	Floor map[string]interface{} `json:"floor"`
	Apartment map[string]interface{} `json:"apartment"`
}

type ConfigBackUrls struct {
	Success map[string]interface{} `json:"success"`
	Pending map[string]interface{} `json:"pending"`
	Failure map[string]interface{} `json:"failure"`
}

type ConfigTaxes struct {
	Typetaxes []TypeTaxesEnum
	Value float64 `json:"value"`
}

type TypeTaxesEnum struct {
	IVA string `json:"IVA"`
	INC string `json:"INC"` 
}

type ModeEnum struct {
	RegularPayment string `json:"regular_payment"`
	MoneyTransfer string `json:"money_transfer"`
}

type AutoReturnEnum struct {
	Approved string `json:"approved"`
	All string `json:"all"`
}