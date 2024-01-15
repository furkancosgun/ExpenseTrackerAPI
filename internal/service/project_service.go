package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type IProjectService interface {
	GetProjectReportByUserId(userId string) ([]dto.ProjectReportDTO, error)
	CreateProject(token dto.ProjectCreateDTO) error
}

type ProjectService struct {
	repository repository.IProjectRepository
}

// CreateProject implements IProjectService.
func (service *ProjectService) CreateProject(dto dto.ProjectCreateDTO) error {
	return service.repository.CreateProject(dto.ToModel())
}

// GetProjectReportByUserId implements IProjectService.
func (service *ProjectService) GetProjectReportByUserId(userId string) ([]dto.ProjectReportDTO, error) {
	return service.repository.GetProjectReportByUserId(userId)
}

func NewProjectService(repository repository.ProjectRepository) IProjectService {
	return &ProjectService{repository: &repository}
}
