package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPostgresqlConnection(ctx context.Context, config Config) *pgxpool.Pool {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnection,
		config.MaxConnectionIdleTime,
	)

	connectionConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		panic(err)
	}

	connection, err := pgxpool.ConnectConfig(ctx, connectionConfig)
	if err != nil {
		panic(err)
	}

	return connection
}
