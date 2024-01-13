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

func NewUserRepository(ctx context.Context, dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{ctx: &ctx, dbPool: dbPool}
}

// GetUserByEmail implements IUserRepository.
func (repository *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	row := repository.dbPool.QueryRow(*repository.ctx, "SELECT * FROM users WHERE lower(email) = $1", email)

	err := row.Scan(&user.Email, &user.FirstName, &user.LastName, &user.Password, &user.AccountConfirmed)

	return user, err
}

// CreateUser implements IUserRepository.
func (repository *UserRepository) CreateUser(user model.User) error {
	user.Normalized()
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO users (email,first_name,last_name,password,account_confirmed) VALUES ($1,$2,$3,$4,$5)",
		user.Email, user.FirstName, user.LastName, user.Password, user.AccountConfirmed,
	)
	return err
}

// DeleteUserByEmail implements IUserRepository.
func (repository *UserRepository) DeleteUserByEmail(email string) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"DELETE users WHERE lower(email) = ?",
		email,
	)
	return err
}

// UpdateUser implements IUserRepository.
func (repository *UserRepository) UpdateUser(user model.User) error {
	user.Normalized()

	_, err := repository.dbPool.Exec(*repository.ctx,
		"UPDATE users SET first_name = $1 ,last_name = $2,password = $3,account_confirmed = $4 WHERE email = $5",
		user.FirstName, user.LastName, user.Password, user.AccountConfirmed, user.Email,
	)
	return err
}
