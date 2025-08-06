package router

import (
	"net/http"

	"vericred/internal/handlers"
	"vericred/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/auth/register", handlers.LoginInMetamask)
		r.Post("/getnonce", handlers.GetNonce)
		r.Post("/auth/metamasklogin", handlers.LoginInMetamask)
	})

	return r
}
