package service

import (
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/furkancosgun/expense-tracker-api/internal/repository"
)

type IExpenseService interface {
	GetExpenses(userId string) ([]model.Expense, error)
	CreateExpense(model model.Expense) error
}

type ExpenseService struct {
	repository repository.IExpenseRepository
}

// CreateExpense implements IExpenseService.
func (service *ExpenseService) CreateExpense(model model.Expense) error {
	return service.repository.CreateExpense(model)
}

// GetExpenses implements IExpenseService.
func (service *ExpenseService) GetExpenses(userId string) ([]model.Expense, error) {
	return service.repository.GetExpenses(userId)
}

func NewExpenseService(repository repository.IExpenseRepository) IExpenseService {
	return &ExpenseService{repository: repository}
}
