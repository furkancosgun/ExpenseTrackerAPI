package repository

import (
	"context"

	"github.com/furkancosgun/expense-tracker-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ITokenRepository interface {
	GetTokenByUserId(userId string) (model.Token, error)
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
		"INSERT INTO tokens (user_id,token,expires_at) VALUES ($1,$2,$3)",
		token.UserId, token.Token, token.ExpiresAt,
	)
	return err
}

// GetTokenByUserId implements ITokenRepository.
func (repository *TokenRepository) GetTokenByUserId(userId string) (model.Token, error) {
	var token model.Token

	row := repository.dbPool.QueryRow(*repository.ctx,
		"SELECT * FROM tokens WHERE user_id = $1", userId,
	)

	err := row.Scan(&token.UserId, &token.Token, &token.ExpiresAt)

	return token, err
}

// UpdateToken implements ITokenRepository.
func (repository *TokenRepository) UpdateToken(token model.Token) error {
	_, err := repository.dbPool.Exec(*repository.ctx,
		"UPDATE tokens SET token = $1 , expires_at = $2 WHERE user_id = $3",
		token.Token, token.ExpiresAt, token.UserId,
	)
	return err
}
