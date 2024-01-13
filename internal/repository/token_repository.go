package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ITokenRepository interface {
	GetTokenByEmail(email string) (model.Token, error)
	CreateToken(token model.Token) error
	UpdateToken(token model.Token) error
}

type TokenRepository struct {
	ctx    *context.Context
	dbPool *pgxpool.Pool
}

func NewTokenRepository(ctx context.Context, dbPool *pgxpool.Pool) ITokenRepository {
	return &TokenRepository{ctx: &ctx, dbPool: dbPool}
}

// CreateToken implements ITokenRepository.
func (repository *TokenRepository) CreateToken(token model.Token) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO tokens (email,token) VALUES ($1,$2)",
		token.Email, token.Token,
	)
	return err
}

// GetTokenByEmail implements ITokenRepository.
func (repository *TokenRepository) GetTokenByEmail(email string) (model.Token, error) {
	var token model.Token

	row := repository.dbPool.QueryRow(*repository.ctx,
		"SELECT * FROM tokens WHERE email = $1",
		email,
	)
	err := row.Scan(&token.Email, &token.Token)

	return token, err
}

// UpdateToken implements ITokenRepository.
func (repository *TokenRepository) UpdateToken(token model.Token) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO tokens (email,token) VALUES ($1,$2)",
		token.Email, token.Token,
	)
	return err
}
