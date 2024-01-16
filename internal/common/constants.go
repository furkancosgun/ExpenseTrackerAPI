package common

import "errors"

const BASE_URL = "api/v1/"

const LOGIN_URL = "/api/v1/auth/login"
const REGISTER_URL = "/api/v1/auth/register"
const FORGOT_PASSWORD_URL = "/api/v1/auth/forgot-password"
const VERIFY_ACCOUNT_URL = "/api/v1/auth/verify-account"
const RESET_PASSWORD_URL = "/api/v1/auth/reset-password"

const CLAIM = "CLAIM"

// Errors
var EMAIL_CANT_BE_EMPTY = errors.New("Email Can't be empty!")
var FIRST_NAME_CANT_BE_EMPTY = errors.New("First Name Can't be empty!")
var LAST_NAME_CANT_BE_EMPTY = errors.New("Last Name Can't be empty!")
var PASSWORD_CANT_BE_EMPTY = errors.New("Password Can't be empty!")
var OTP_CODE_CANT_BE_EMPTY = errors.New("OTP Code Can't be empty!")
var EMAIL_ALREADY_USING = errors.New("Email Already Using!")
var USER_NOT_FOUND = errors.New("User Not Found!")
var UN_CONFIRMED_ACCOUNT = errors.New("Account Not Confirmed!")
var INVALID_OTP_TOKEN = errors.New("Invalid OTP Token")
var USER_ID_CANT_BE_EMPTY = errors.New("UserId Can't be empty")
var CATEGORY_NAME_CANT_BE_EMPTY = errors.New("Category Name Can't be empty")

var NOT_REQUIRED_AUTH_CHECK_URLS = []string{
	LOGIN_URL,
	REGISTER_URL,
	FORGOT_PASSWORD_URL,
	VERIFY_ACCOUNT_URL,
	RESET_PASSWORD_URL,
}

const JWT_KEY = "VERY_IMPORTANT_JWT_AUTH_KEY"
