package database

import "os"

type Config struct {
	Host                  string
	Port                  string
	UserName              string
	Password              string
	DbName                string
	MaxConnection         string
	MaxConnectionIdleTime string
}

func NewConfig() *Config {
	return &Config{
		Host:                  os.Getenv("PSQL_HOST"),
		Port:                  os.Getenv("PSQL_PORT"),
		DbName:                os.Getenv("PSQL_DBNAME"),
		UserName:              os.Getenv("PSQL_USERNAME"),
		Password:              os.Getenv("PSQL_PASSWORD"),
		MaxConnection:         os.Getenv("PSQL_MAXCONN"),
		MaxConnectionIdleTime: os.Getenv("PSQL_MAXCONNIDLETIME"),
	}
}
