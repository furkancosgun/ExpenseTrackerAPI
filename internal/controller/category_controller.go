package controller

import (
	"encoding/json"
	"net/http"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (controller *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value(common.CLAIM).(common.Claim)

	categories, err := controller.service.GetCategories(claim.UserId)

	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	helper.JsonWriteToResponse(w, categories, http.StatusOK)
}
func (controller *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	claim := r.Context().Value(common.CLAIM).(common.Claim)

	var dto dto.CategoryCreateDTO

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = dto.Validate()
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	dto.UserId = claim.UserId

	err = controller.service.CreateCategory(dto)
}
