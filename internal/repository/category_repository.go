package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ICategoryRepository interface {
	GetCategories(userId string) ([]model.Category, error)
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
		"INSERT INTO categories (id,user_id,name) VALUES ($1,$2,$3)",
		category.Id, category.UserId, category.Name,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *CategoryRepository) GetCategories(userId string) ([]model.Category, error) {
	var categories []model.Category

	row, err := repository.dbPool.Query(*repository.ctx,
		"SELECT * FROM categories WHERE user_id = $1", userId,
	)

	var category model.Category
	for row.Next() {
		err = row.Scan(&category.Id, &category.UserId, &category.Name)
		if err != nil {
			break
		}
		categories = append(categories, category)
	}

	return categories, err
}
