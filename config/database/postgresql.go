package database

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
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "expense_tracker",
		UserName:              "postgres",
		Password:              "postgres",
		MaxConnection:         "10",
		MaxConnectionIdleTime: "100s",
	}
}
