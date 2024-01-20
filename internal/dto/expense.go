package dto

type CreateExpenseRequest struct {
	ProjectId    string
	MerchantName string
	Amount       float32
	Date         string
	Description  string
	CategoryId   string
	IncludeVat   bool
	Vat          float32
}
