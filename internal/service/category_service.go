package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type ICategoryService interface {
	GetCategories(userId string) ([]dto.ListCategoryResponse, error)
	CreateCategory(model model.Category) error
}

type CategoryService struct {
	repository repository.ICategoryRepository
}

// CreateCategory implements ICategoryService.
func (service *CategoryService) CreateCategory(model model.Category) error {
	return service.repository.CreateCategory(model)
}

// GetCategories implements ICategoryService.
func (service *CategoryService) GetCategories(userId string) ([]dto.ListCategoryResponse, error) {
	return service.repository.GetCategories(userId)
}

func NewCategoryService(repository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{repository: repository}
}
