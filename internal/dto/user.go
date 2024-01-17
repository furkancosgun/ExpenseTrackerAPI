package dto

type UserLoginRequest struct {
	Email    string
	Password string
}

type UserResetPasswordRequest struct {
	Email    string
	Password string
	Otp      string
}

type UserVerifyAccountRequest struct {
	Email string
	Otp   string
}
type UserRegisterRequest struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type UserLoginResponse struct {
	FirstName string
	LastName  string
	Email     string
	Token     string
}

type UserForgotPasswordRequest struct {
	Email string
}
