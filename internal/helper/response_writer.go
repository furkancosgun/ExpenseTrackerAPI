package helper

import (
	"encoding/json"
	"net/http"

	"github.com/furkancosgun/expense-tracker-api/internal/common"
)

func JsonWriteToErrorResponse(w http.ResponseWriter, value error, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(common.NewErrorResponseFromError(value))
}
func JsonWriteToResponse(w http.ResponseWriter, value any, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(value)
}
