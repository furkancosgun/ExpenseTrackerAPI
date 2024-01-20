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

type ExpenseController struct {
	service service.IExpenseService
}

func NewExpenseController(service service.IExpenseService) *ExpenseController {
	return &ExpenseController{service: service}
}

func (controller *ExpenseController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	categories, err := controller.service.GetExpenses(userId)

	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	helper.JsonWriteToResponse(w, categories, http.StatusOK)
}
func (controller *ExpenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(common.AUTH_USER_ID).(string)

	var dto dto.CreateExpenseRequest

	//Json Decode
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	model := model.Expense{
		ExpenseId:    uuid.NewString(),
		ProjectId:    dto.ProjectId,
		UserId:       userId,
		MerchantName: dto.MerchantName,
		Amount:       dto.Amount,
		Date:         dto.Date,
		Description:  dto.Description,
		CategoryId:   dto.CategoryId,
		IncludeVat:   dto.IncludeVat,
		Vat:          dto.Vat,
	}

	err = controller.service.CreateExpense(model)
	if err != nil {
		helper.JsonWriteToErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
