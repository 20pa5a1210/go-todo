package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, statuscode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	errResp := ErrorResponse{
		Error: errorMessage,
	}
	json.NewEncoder(w).Encode(errResp)
}
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
