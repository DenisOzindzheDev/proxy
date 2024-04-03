package responses

import (
	"encoding/json"
	"net/http"
)

// Writes OK responses
func StatusOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

// Writes Error responses
func StatusError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"header": "Application Error",
		"error":  message,
	}
	json.NewEncoder(w).Encode(body)
}
