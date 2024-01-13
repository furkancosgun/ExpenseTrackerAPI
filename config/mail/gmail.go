package mail

import "os"

type Config struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     string
}

func NewCofig() *Config {
	return &Config{
		Identity: os.Getenv("SMTP_ID"),
		Username: os.Getenv("SMTP_ADDR"),
		Password: os.Getenv("SMTP_KEY"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}
}
