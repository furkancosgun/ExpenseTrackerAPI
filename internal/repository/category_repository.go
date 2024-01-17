package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/dto"
	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ICategoryRepository interface {
	GetCategories(userId string) ([]dto.ListCategoryResponse, error)
	CreateCategory(token model.Category) error
}

type CategoryRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewCategoryRepository(ctx context.Context, dbPool *pgxpool.Pool) ICategoryRepository {
	return &CategoryRepository{ctx: &ctx, dbPool: dbPool}
}

// CreateToken implements ITokenRepository.
func (repository *CategoryRepository) CreateCategory(category model.Category) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO categories (category_id,user_id,name) VALUES ($1,$2,$3)",
		category.CategoryId, category.UserId, category.Name,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *CategoryRepository) GetCategories(userId string) ([]dto.ListCategoryResponse, error) {
	var categories []dto.ListCategoryResponse

	row, err := repository.dbPool.Query(*repository.ctx,
		"SELECT category_id,name FROM categories WHERE user_id = $1", userId,
	)

	var category dto.ListCategoryResponse
	for row.Next() {
		err = row.Scan(&category.CategoryId, &category.Name)
		if err != nil {
			break
		}
		categories = append(categories, category)
	}

	return categories, err
}
