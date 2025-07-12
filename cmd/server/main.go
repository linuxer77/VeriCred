package main

import (
	"fmt"
	"net/http"
	"vericred/internal/eth"

	"github.com/go-chi/chi/v5"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This shit is working now"))
}

func main() {
	eth.Contractresp("BODY")
	
	r := chi.NewRouter()
	r.Get("/", handleIndex)

	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
