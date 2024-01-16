package dto

import (
	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
)

type CategoryCreateDTO struct {
	UserId string
	Name   string
}

func (dto *CategoryCreateDTO) ToModel() model.Category {
	return model.Category{UserId: dto.UserId, Name: dto.Name}
}

func (dto *CategoryCreateDTO) Validate() error {
	if dto.Name == "" {
		return common.CATEGORY_NAME_CANT_BE_EMPTY
	}
	return nil
}
