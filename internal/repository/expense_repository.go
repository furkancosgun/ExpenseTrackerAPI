package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IExpenseRepository interface {
	GetExpenses(userId string) ([]model.Expense, error)
	CreateExpense(token model.Expense) error
}

type ExpenseRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewExpenseRepository(ctx context.Context, dbPool *pgxpool.Pool) IExpenseRepository {
	return &ExpenseRepository{ctx: &ctx, dbPool: dbPool}
}

// CreateToken implements ITokenRepository.
func (repository *ExpenseRepository) CreateExpense(expense model.Expense) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO expenses (id,project_id,merchant_name,amount,date,description,category_id,include_vat,vat,image_path) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		expense.Id, expense.ProjectId, expense.MerchantName, expense.Amount, expense.Date, expense.Description, expense.CategoryId, expense.IncludeVat, expense.Vat, expense.ImagePath,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *ExpenseRepository) GetExpenses(userId string) ([]model.Expense, error) {
	var expenses []model.Expense

	row, err := repository.dbPool.Query(*repository.ctx,
		"SELECT * FROM expenses WHERE user_id = $1", userId,
	)

	var expense model.Expense
	for row.Next() {
		err = row.Scan(&expense.Id, &expense.ProjectId, &expense.MerchantName, &expense.Amount, &expense.Date, &expense.Description, &expense.CategoryId, &expense.IncludeVat, &expense.Vat, &expense.ImagePath)
		if err != nil {
			break
		}
		expenses = append(expenses, expense)
	}

	return expenses, err
}
