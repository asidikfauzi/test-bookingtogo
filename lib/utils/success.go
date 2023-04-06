package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func StatusOK(w http.ResponseWriter, code int, message string, data interface{}) {
	response := map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.New("Failed to encode response")
	}
}
