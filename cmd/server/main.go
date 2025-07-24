package main

import (
	"fmt"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/router"
)

// func handleIndex(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("This shit is working now"))
// }

func main() {
	db.Init()	
	r := router.RegisterRouter()
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
