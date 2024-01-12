package common

const BASE_URL = "api/v1/"

const LOGIN_URL = "api/v1/auth/login"
const REGISTER_URL = "api/v1/auth/register"
const FORGOT_PASSWORD_URL = "api/v1/auth/forgot-password"
const VERIFY_ACCOUNT_URL = "api/b1/auth/verify-account"

var NOT_REQUIRED_AUTH_CHECK_URLS = []string{
	LOGIN_URL,
	REGISTER_URL,
	FORGOT_PASSWORD_URL,
	VERIFY_ACCOUNT_URL,
}

const JWT_KEY = "VERY_IMPORTANT_JWT_AUTH_KEY"
