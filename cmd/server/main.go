package main

import (
	"fmt"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/router"
)

func main() {
	db.Init()	
	r := router.RegisterRouter()
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
