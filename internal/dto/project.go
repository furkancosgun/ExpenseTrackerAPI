package dto

import (
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
)

type ProjectCreateDTO struct {
	UserId string
	Name   string
}

func (dto *ProjectCreateDTO) ToModel() model.Project {
	return model.Project{UserId: dto.UserId, Name: dto.Name, CreatedAt: time.Now()}
}

type ProjectReportDTO struct {
	MerchantName  string
	TotalAmount   float64
	CreatedAt     time.Time
	TotalExpenses int8
}
