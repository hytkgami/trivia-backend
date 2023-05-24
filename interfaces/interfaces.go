package interfaces

import (
	"encoding/json"
	"net/http"
)

func HttpErrorResponse(w http.ResponseWriter, err error, status int) {
	msg := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(msg)
}
