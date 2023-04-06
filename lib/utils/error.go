package utils

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"net/http"
	"test-bookingtogo/models"
)

type ErrorMessageEmpty struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMessageEmpty(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}
	return "Unknown error"
}

func BadRequestErrorFieldEmpty(w http.ResponseWriter, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMessageEmpty, len(ve))
		for i, e := range err.(validator.ValidationErrors) {
			out[i] = ErrorMessageEmpty{
				Field:   e.Field(),
				Message: GetErrorMessageEmpty(e),
			}
		}
		BadRequest(w, out, "Bad Request")
	}
	return
}

func BadRequest(w http.ResponseWriter, message interface{}, status string) {
	var errorResponse = models.ErrorResponse{
		http.StatusBadRequest,
		message,
		status,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(errorResponse)
}

func InternalServerError(w http.ResponseWriter, message interface{}) {
	var errorResponse = models.ErrorResponse{
		http.StatusInternalServerError,
		message,
		"Internal Server Error",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorResponse)
}
