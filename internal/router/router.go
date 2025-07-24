package router

import (
	"net/http"

	"vericred/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/auth/register", handlers.LoginInMetamask)
	return r
}
