package service

import (
	"time"

	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type IProjectService interface {
	GetProjectReportByUserId(userId string) ([]dto.ProjectReportResponse, error)
	CreateProject(model model.Project) error
	GetProjects(userId string) ([]dto.ProjectListResponse, error)
}

type ProjectService struct {
	repository repository.IProjectRepository
}

// GetProjects implements IProjectService.
func (service *ProjectService) GetProjects(userId string) ([]dto.ProjectListResponse, error) {
	return service.repository.GetProjects(userId)
}

// CreateProject implements IProjectService.
func (service *ProjectService) CreateProject(model model.Project) error {
	model.CreatedAt = time.Now()
	return service.repository.CreateProject(model)
}

// GetProjectReportByUserId implements IProjectService.
func (service *ProjectService) GetProjectReportByUserId(userId string) ([]dto.ProjectReportResponse, error) {
	return service.repository.GetProjectReportByUserId(userId)
}

func NewProjectService(repository repository.IProjectRepository) IProjectService {
	return &ProjectService{repository: repository}
}
