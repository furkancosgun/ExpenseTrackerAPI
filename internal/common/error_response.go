package common

type ErrorResponse struct {
	ErrorDescription string
}

func NewErrorResponseFromError(err error) ErrorResponse {
	return ErrorResponse{ErrorDescription: err.Error()}
}

func NewErrorResponseFromString(err string) ErrorResponse {
	return ErrorResponse{ErrorDescription: err}
}
