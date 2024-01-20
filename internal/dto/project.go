package dto

import (
	"time"
)

type CreateProjectRequest struct {
	Name string
}

type ProjectReportResponse struct {
	ProjectId     string
	ProjectName   string
	TotalAmount   float64
	CreatedAt     time.Time
	TotalExpenses int8
}

type ProjectListResponse struct {
	ProjectId   string
	ProjectName string
}
