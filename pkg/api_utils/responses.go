package responses

import (
	"encoding/json"
	"net/http"
)

// Writes OK responses
func StatusOk(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
}

// Writes Error responses
func StatusError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	body := map[string]string{
		"header": "Application Error",
		"error":  message,
	}
	json.NewEncoder(w).Encode(body)
}
