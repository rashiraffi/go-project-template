package handler

import (
	"net/http"
	"template/service"
)

type handler struct {
	service service.Service
}

type Handler interface {
	// Add new methods here
	Login(w http.ResponseWriter, r *http.Request)
}

func NewHandler(
	service service.Service,
) Handler {
	return &handler{
		service: service,
	}
}
