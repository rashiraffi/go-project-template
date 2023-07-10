package router

import (
	"template/handler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter(handler handler.Handler) (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", handler.Login)

	return r
}
