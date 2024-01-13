package test

import (
	"testing"

	"github.com/furkancosgun/expense-tracker-api/internal/helper"
)

func Test(t *testing.T) {

	t.Run("OTP Generato Test", func(t *testing.T) {
		var otp = helper.GenerateOTP(6)
		if len(otp) != 6 {
			t.Error("Lenght Not Mached")
		}
		t.Log(otp)
	})
}
