package model
import ()

type MessageResponse struct {
	Error   string `json:"error"`   
	Message string `json:"message"` 
	Status  int    `json:"status"`
}

type PaymentResponse struct {
	UserId int `json:"user_id"`
	StatusDetails string `json:"status_details"`
	OperationId string `json:"operation_id"`
	DateApproved string `json:"date_approved"`
	Payer []string `json:"payer"`
	PaymentMethods PaymentMethod `json:"payment_methods"`
	PaymentTypeId string `json:"payment_type_id"`
	Refunds string `json:"refunds"`
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