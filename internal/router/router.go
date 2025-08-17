package router

import (
	"net/http"

	"vericred/internal/handlers"
	"vericred/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func RegisterRouter() http.Handler {
	r := chi.NewRouter()
    r.Use(middleware.CORSMiddleware)
	r.Use(middleware.LoggingMiddleware)
	r.Post("/getnonce", handlers.GetNonce)
	r.Post("/auth/metamasklogin", handlers.LoginInMetamask)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/api/create/user", handlers.CreateUser)
		r.Post("/api/create/org", handlers.CreateUniversity)
		r.Get("/dashboard", handlers.ShowUser)
		r.Get("/university", handlers.ShowOrg)
		r.Get("/universities", handlers.AllOrgs)
		// r.Get("/university", handlers.ShowUniversity)
	})
	return r
}
