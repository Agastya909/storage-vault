package utils

import (
	"encoding/json"
	"net/http"
)

func HttpResponseHandler(w http.ResponseWriter, statusCode int, message string, data, headers any) {
	w.Header().Add("Content-Type", "application/json")
	if headers != nil {
		for key, value := range headers.(map[string]string) {
			w.Header().Set(key, value)
		}
	}

	response := map[string]any{
		"message": message,
		"data":    data,
	}
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
