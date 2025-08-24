package main

import (
	"fmt"
	_ "math/big"
	"net/http"

	// "vericred/internal/db"
	// "vericred/internal/logging"
	"vericred/internal/db"
	"vericred/internal/logging"
	"vericred/internal/router"
)

func main() {
	db.Init()	
	logging.Init()

	r := router.RegisterRouter()
	
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
