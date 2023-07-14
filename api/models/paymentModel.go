package model

import ()

type PayerModel struct {
	Properties PropertiesPayer
	BinaryMode bool `json:"boolean"`
	Order PropertiesOrder
	ExternalReference string `json:"external_reference"`
	Description string `json:"description"`
	Metadata map[string]interface{}
	TransactionAmount int `json:"transaction_amount"`
	CouponAmount int `json:"coupon_amount"`
	CampaignId int `json:"campaign_id"`
	CouponCode string `json:"coupon_code"`
	DifferentialPricingId int `json:"differential_pricing_id"`
	ApplicationFee int `json:"application_fee"`
	Capture bool `json:"capture"`
	PaymentMethodId string `json:"payment_method_id"`
	IssuerId string `json:"issuer_id"`
	Token string `json:"token"`
	StatementDescriptor string `json:"statement_descriptor"`
	Installments int `json:"installments"`
	NotificationUrl string `json:"notification_url"`
	CallbackUrl string `json:"callback_url"`
	AdditionalInfo AdditionalInfo
	TransactionDetails TransactionDetails
}

type TransactionDetails struct {
	Transaction TransactionProp
}

type TransactionProp struct {
	AcquirerReference string `json:"acquirer_reference"`
	BankTransferId int `json:"bank_transfer_id"`
	ExternalResourceUrl string `json:"external_resource_url"`
	FinancialInstitution string `json:"financial_institution"`
	InstallmentAmount int `json:"installment_amount"`
	NetReceivedAmount int `json:"net_received_amount"`
	OverpaidAmount int `json:"overpaid_amount"`
	PayableDeferralPeriod string `json:"payable_deferral_period"`
	PaymentMethodReferenceId string `json:"payment_method_reference_id"`
	TotalPaidAmount int `json:"total_paid_amount"`
	TransactionCicle string `json:"transaction_cicle"`
}

type AdditionalInfo struct {
	Properties PropAddInfo
}

type PropAddInfo struct {
	IpAddress string `json:"ip_address"`
	Items Items
	Payer Payer
	Shipments ShippingProp
}

type ShippingProp struct {
	ZipCode string `json:"zip_code"`
	StreetName string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Floor int `json:"floor"`
	Apartment string `json:"apartment"`
}

type Payer struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Phone PhonePayer
	Address AddressPayer
	RegistrationDate string `json:"registration_date"`
}

type PhonePayer struct {
	AreaCode string `json:"areaCode"`
	Number string `json:"number"`
}

type AddressPayer struct {
	ZipCode string `json:"zipCode"`
	StreetName string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
}

type Items struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	PictureUrl string `json:"picture_url"`
	CategoryId string `json:"category_id"`
	Quantity int `json:"quantity"`
	UnitPrice int `json:"unit_price"`
}

type PropertiesOrder struct {
	MercadoLibre string `json:"mercadolibre"`
	MercadoPago string `json:"mercadopago"`
}

type PropertiesPayer struct {
	EntityType EntityTypeEnum
	Type TypeEnum
	Id string `json:"id"`
	Email Email
	Identification Identification
	Phone Phone
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type EntityTypeEnum struct {
	Individual string `json:"individual"`
	Association string `json:"association"`
}

type TypeEnum struct {
	Customer string `json:"customer"`
	Registered string `json:"registered"`
	Guest string `json:"guest"`
}

type Email struct {
	Type string `json:"type"`
	Format string `json:"email"`
}

type Identification struct {
	Type string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	AreaCode string `json:"areaCode"`
	Number string `json:"number"`
	Extension string `json:"extension"`
}
