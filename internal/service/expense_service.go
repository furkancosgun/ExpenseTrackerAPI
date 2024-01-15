package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type IExpenseService interface {
	GetExpenses(userId string) ([]model.Expense, error)
	GetProjectByUserId(userId string) (model.Expense, error)
	CreateExpense(token dto.CreateExpenseDTO) error
}

type ExpenseService struct {
	repository repository.IExpenseRepository
}

// CreateExpense implements IExpenseService.
func (service *ExpenseService) CreateExpense(dto dto.CreateExpenseDTO) error {
	return service.repository.CreateExpense(dto.ToModel())
}

// GetExpenses implements IExpenseService.
func (service *ExpenseService) GetExpenses(userId string) ([]model.Expense, error) {
	return service.repository.GetExpenses(userId)
}

// GetProjectByUserId implements IExpenseService.
func (service *ExpenseService) GetProjectByUserId(userId string) (model.Expense, error) {
	return service.repository.GetProjectByUserId(userId)
}

func NewExpenseService(repository repository.ExpenseRepository) IExpenseService {
	return &ExpenseService{repository: &repository}
}
