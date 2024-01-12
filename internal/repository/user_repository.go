package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IUserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	DeleteUserByEmail(email string) error
}

type UserRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}
