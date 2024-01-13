package test

import (
	"testing"

	"github.com/furkancosgun/expense-tracker-api/internal/service"
	"github.com/joho/godotenv"
)

func TestMailService(t *testing.T) {
	t.Run("SEND", func(t *testing.T) {

		godotenv.Load("../.env")
		mail := service.NewOtpMailContent("rehod94758@talmetry.com", "123")
		err := mail.Send()
		if err != nil {
			t.FailNow()
		}
	})
}
