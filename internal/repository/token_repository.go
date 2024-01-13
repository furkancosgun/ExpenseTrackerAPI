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
	token.Normalized()
	_, err := repository.dbPool.Exec(*repository.ctx,
		"INSERT INTO tokens (email,token,expires_at) VALUES ($1,$2,$3)",
		token.Email, token.Token, token.ExpiresAt,
	)
	return err
}

// GetTokenByEmail implements ITokenRepository.
func (repository *TokenRepository) GetTokenByEmail(email string) (model.Token, error) {
	var token model.Token

	row := repository.dbPool.QueryRow(*repository.ctx,
		"SELECT * FROM tokens WHERE lower(email) = $1",
		email,
	)

	err := row.Scan(&token.Email, &token.Token, &token.ExpiresAt)

	return token, err
}

// UpdateToken implements ITokenRepository.
func (repository *TokenRepository) UpdateToken(token model.Token) error {
	token.Normalized()
	_, err := repository.dbPool.Exec(*repository.ctx,
		"UPDATE tokens SET token = $1 , expires_at = $2 WHERE lower(email) = $3",
		token.Token, token.ExpiresAt, token.Email,
	)
	return err
}
