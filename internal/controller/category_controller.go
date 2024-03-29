package controller

import (
	"encoding/json"
	"net/http"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/google/uuid"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (controller *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	categories, err := controller.service.GetCategories(userId)

	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	helper.JsonWriteToResponse(w, categories, http.StatusOK)
}
func (controller *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	var dto dto.CreateCategoryRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	model := model.Category{
		CategoryId: uuid.New().String(),
		Name:       dto.Name,
		UserId:     userId,
	}

	err = controller.service.CreateCategory(model)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
