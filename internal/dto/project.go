package dto

import (
	"time"
)

type CreateProjectRequest struct {
	Name string
}

type ListProjectResponse struct {
	ProjectId     string
	ProjectName   string
	TotalAmount   float64
	CreatedAt     time.Time
	TotalExpenses int8
}
