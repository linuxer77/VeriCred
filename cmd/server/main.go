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
	C := eth.Contract{}
	// Contract address for now: 0xDE5C084a7959533893954BA072895B53fE1E7486
	// Current verified orgs: 0x2e176D534611f09dEa9BB9DB02a23Aa982F60ec9
	// Owner address: 0x10F315866210212BAee0c5A45d6D2a5d772d7ceD 
	C.InteractViewFunction("0xDE5C084a7959533893954BA072895B53fE1E7486", "")
	// C.InteractTransactionFunction("0xDE5C084a7959533893954BA072895B53fE1E7486", "0x7cB57B5A97eAbe94205C07890BE4c1b59C04f6b7")

	r := chi.NewRouter()
	r.Get("/", handleIndex)

	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
