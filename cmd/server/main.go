package main

import (
	"fmt"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/handlers"

	"github.com/go-chi/chi/v5"
)

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This shit is working now"))
// }

func main() {
	db.Init()	
	r := chi.NewRouter()
	r.Post("/auth/user/register", handlers.CreateUser)
	r.Post("/auth/org/register", handlers.CreateOrg)

	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
