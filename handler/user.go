package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"template/entities"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		log.Default().Println("Empty body")
		sendResponse(w, http.StatusBadRequest, "Empty body")
		return
	}

	body := entities.User{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, "Invalid body")
		return
	}

	token, err := h.service.Login(&body)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if token == "" {
		sendResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	sendResponse(
		w,
		http.StatusOK,
		map[string]string{
			"token": token,
		},
	)
}
