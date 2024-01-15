package model

import "time"

type Expense struct {
	Id           string
	ProjectId    string
	MerchantName string
	Amount       float32
	Date         time.Time
	Description  string
	CategoryId   string
	IncludeVat   bool
	Vat          float32
	ImagePath    string
}
