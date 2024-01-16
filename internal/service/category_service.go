package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
	"github.com/google/uuid"
)

type ICategoryService interface {
	GetCategories(userId string) ([]model.Category, error)
	CreateCategory(dto dto.CategoryCreateDTO) error
}

type CategoryService struct {
	repository repository.ICategoryRepository
}

// CreateCategory implements ICategoryService.
func (service *CategoryService) CreateCategory(dto dto.CategoryCreateDTO) error {
	err := dto.Validate()
	if err != nil {
		return err
	}
	model := dto.ToModel()

	model.Id = uuid.New().String()

	return service.repository.CreateCategory(model)
}

// GetCategories implements ICategoryService.
func (service *CategoryService) GetCategories(userId string) ([]model.Category, error) {
	return service.GetCategories(userId)
}

func NewCategoryService(repository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{repository: repository}
}
