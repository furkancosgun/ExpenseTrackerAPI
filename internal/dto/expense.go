package dto

import (
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
)

type CreateExpenseDTO struct {
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

func (dto *CreateExpenseDTO) ToModel() model.Expense {
	return model.Expense{
		ProjectId:    dto.ProjectId,
		MerchantName: dto.MerchantName,
		Amount:       dto.Amount,
		Date:         dto.Date,
		Description:  dto.Description,
		CategoryId:   dto.CategoryId,
		IncludeVat:   dto.IncludeVat,
		Vat:          dto.Vat,
		ImagePath:    dto.ImagePath,
	}
}
