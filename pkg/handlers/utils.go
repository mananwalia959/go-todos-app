package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mananwalia959/go-todos-app/pkg/models"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, 404, "Path Not Found")
}

func encodeToJson(w http.ResponseWriter, status int, jsonInterface interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonInterface)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	errorMessage := models.ErrorResponse{Message: message}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage)
}
