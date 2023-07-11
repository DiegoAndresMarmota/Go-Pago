package model
import ()

type MessageResponse struct {
	Error   string `json:"error"`   
	Message string `json:"message"` 
	Status  int    `json:"status"`
}

type PaymentResponse struct {
	UserId int `json:"user_id"`
	DateCreated string `json:"date_created"`
	DateApproved string `json:"date_approved"`
	DateLastUpdated string `json:"date_last_updated"`
	DateExpiration string `json:"date_of_expiration"`
	MoneyReleaseDate string `json:"money_release_date"`
	OperationType string `json:"operation_type"`
	PaymentMethodId string `json:"payment_method_id"`
	PaymentTypeId string `json:"payment_type_id"`
	Status string `json:"status"`
	StatusDetail string `json:"status_detail"`
	CurrencyId string `json:"currency_id"`
	Description string `json:"description"`


	StatusDetails string `json:"status_details"`
	OperationId string `json:"operation_id"`
	Payer []string `json:"payer"`
	PaymentMethods PaymentMethod `json:"payment_methods"`
	Refunds string `json:"refunds"`
}

type PaymentOverload struct {
	Id string `json:"id"`
	Payments []Payments `json:"payments"`
	Currency string `json:"currency"`
	Amount int `json:"amount"`
	Reason string `json:"reason"`
	CoverageApplied bool `json:"coverage_applied"`
	CoverageElegible bool `json:"coverage_eligible"`
	DocumentationRequired bool `json:"documentation_required"`
	DocumentationStatus string `json:"documentation_status"`
	Documentation Documentation `json:"documentation"`
	DateDocumentationDeadline string `json:"date_documentation_deadline"`
	DateCreated string `json:"date_created"`
	DateLastUpdated string `json:"date_last_updated"`
	LiveMode bool `json:"live_mode"`

}




type PaymentRequest struct {

}

type Payer struct {
	Payer PayerId `json:"payer_id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Address PayerAddress `json:"address"`
	Phone PayerPhone `json:"phone"`
}

type PayerId struct {
	Identification string `json:"identification"`
	Number int `json:"number"`
}

type PayerPhone struct {
	AreaCodeLocation int `json:"area_code_location"`
	Number int `json:"number"`
}

type PayerAddress struct {
	Name string `json:"name_street"`
	Number int `json:"number_street"`
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
	ZipCode int `json:"zip_code"`
}

type PaymentMethod struct {

}

type Payments struct {

}

type Documentation struct {
	Type string `json:"type"`
	Url string `json:"url"`
	Description string `json:"description"`

}