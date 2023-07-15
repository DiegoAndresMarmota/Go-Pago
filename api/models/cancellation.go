package models

import ()

type Cancellation struct {
	Id int `json:"id"`
	DateCreated	string `json:"dateCreated"`
	DateApproved string `json:"dateApproved"`
	DateLastUpdated string `json:"date_last_Updated"`
	DateExpiration string `json:"date_of_expiration"`
	MoneyReleaseDate string `json:"money_release_date"`
	OperationType string `json:"operation_type"`
	IssuerId int `json:"issuer_id"`
	PaymentMethodId int `json:"payment_method_id"`
	Status string `json:"status"`
	StatusDetail string `json:"status_detail"`
	CurrencyId string `json:"currency_id"`
	Description string `json:"description"`
	LiveMode bool `json:"live_mode"`
	SponsorId int `json:"sponsor_id"`
	AuthorizationCode string `json:"authorization_code"`
	MoneyReleaseSchema string `json:"money_release_schema"`
	TaxesAmount int `json:"taxes_amount"`
	CounterCurrency string `json:"counter_currency"`
	BrandId int `json:"brand_id"`
	ShippingAmount int `json:"shipping_amount"`
	PosId int `json:"pos_id"`
	StoreId int `json:"store_id"`
	IntegratorId int `json:"integrator_id"`
	PlatformId int `json:"platform_id"`
	CorporationId int `json:"corporatoin_id"`
	CollectorIde int `json:"collector_id"`
	Payers Payers `json:"payers"`
	MarketplaceOwner string `json:"marketplace_owner"`
	Metadata map[string]interface{}
	AdditionalInfo AdditionalInfoCancel
	Order Order `json:"order"`
	ExternalReference string `json:"external_reference"`
	TransactionAmount float64 `json:"transaction_amount"`
	CouponAmount int `json:"coupon_amount"`
	DifferentialPricingId int `json:"differential_pricing_id"`
	DeductionSchema string `json:"deduction_schema"`
	Barcode Barcode `json:"barcode"`
	Installments int `json:"installments"`
	TransactionDetails TransactionDetailsCancel
	FeeDetails map[string]interface{} `json:"fee_details"`
	ChargesDetails map[string]interface{} `json:"charges_details"`
	Captured bool `json:"captured"`
	BinaryMode bool `json:"binar_mode"`
	CallForAuthorizaId string `json:"call_for_authorization_id"`
	StatementDescriptor string `json:"statement_descriptor"`
	Card map[string]interface{} `json:"card"`
	NotificationUrl string `json:"notification_url"`
	Refunds map[string]interface{} `json:"refunds"`
	ProcessingMode string `json:"processing_mode"`
	MerchantAccountId int `json:"merchant_account_id"`
	MerchantNumber string `json:"merchant_number"`
	AcquirerReconciliation map[string]interface{} `json:"acquirer_reconciliation"`
	PointOfInteraction PointOfInteraction
}

type Payers struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Type string `json:"type"`
	Identification IdentificationPayers
	Phone PhonePayers `json:"phone"`
	EntityType string `json:"entity_type"`
	EntityId int `json:"entity_id"`
	OperatorId int `json:"operator_id"`
}

type IdentificationPayers struct {
	Number int `json:"number"`
	Type string `json:"type"`
}

type PhonePayers struct {
	AreaCode int `json:"area_code"`
	Number int `json:"number"`
	Extension int `json:"extension"`
}

type AdditionalInfoCancel struct {
	Items ItemsAddInfo `json:"items"`
}

type ItemsAddInfo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	PictureUrl string `json:"picture_url"`
	CategoryId string `json:"category_id"`
	Quantity int `json:"quantity"`
	UnitPrice int `json:"unit_price"`
}

type Order struct {
	Type string `json:"type"`
	Id int `json:"id"`
}

type Barcode struct {
	Content string `json:"content"`
}

type TransactionDetailsCancel struct {
	PaymentMethodReferenceId string `json:"payment_method_reference_id"`
	VerificationCode int `json:"verification_code"`
	NetReceivedCode int `json:"net_received_code"`
	TotalPaidAmount float64 `json:"total_paid_amount"`
	OverPaidAmount int `json:"over_paid_amount"`
	ExternalResourceUrl string `json:"external_resource_url"`
	InstallmentAmount int `json:"installment_amount"`
	FinancialInstitution string `json:"financial_institution"`
	PayableDeferalPeriod int `json:"payable_deferal_period"`
	AcquirerReference string `json:"acquirer_reference"`
}

type PointOfInteraction struct {
	Type string `json:"type"`
	BusinessInfo BusinessInfo
}

type BusinessInfo struct {
	Unit string `json:"unit"`
	SubUnit string `json:"sub_unit"`
}