package handler

import (
	"encoding/json"
	"net/http"
)

// sendResponse is a helper function to send a response to the client
func sendResponse(w http.ResponseWriter, statusCode int, payload any) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Write(response) // nolint: errcheck
	return
}
