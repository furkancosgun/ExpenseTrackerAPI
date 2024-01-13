package test

import (
	"testing"

	"github.com/furkancosgun/expense-tracker-api/internal/helper"
	"github.com/joho/godotenv"
)

func TestMailService(t *testing.T) {
	t.Run("SEND", func(t *testing.T) {

		godotenv.Load("../.env")
		mail := helper.NewMailContent("rehod94758@talmetry.com", "deneme", "DENEME CONTENT")
		err := mail.Send()
		if err != nil {
			t.FailNow()
		}
	})
}
