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
}

type UserRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewUserRepository(ctx context.Context, dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{ctx: &ctx, dbPool: dbPool}
}

// GetUserByEmail implements IUserRepository.
func (repository *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	row := repository.dbPool.QueryRow(*repository.ctx, "SELECT * FROM users WHERE email = $1", email)

	err := row.Scan(&user.UserId, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.AccountConfirmed)

	return user, err
}

// CreateUser implements IUserRepository.
func (repository *UserRepository) CreateUser(user model.User) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO users (user_id,email,first_name,last_name,password,account_confirmed) VALUES ($1,$2,$3,$4,$5,$6)",
		user.UserId, user.Email, user.FirstName, user.LastName, user.Password, user.AccountConfirmed,
	)
	return err
}

// UpdateUser implements IUserRepository.
func (repository *UserRepository) UpdateUser(user model.User) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"UPDATE users SET first_name = $1 ,last_name = $2,password = $3,account_confirmed = $4 email = $5 WHERE user_id = $6",
		user.FirstName, user.LastName, user.Password, user.AccountConfirmed, user.Email, user.UserId,
	)
	return err
}
