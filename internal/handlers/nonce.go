package handlers

import (
	"encoding/json"
	"net/http"
	"vericred/pkg"
)

var nonceMap = map[string]string{}

func GetNonce(w http.ResponseWriter, r *http.Request) {
	var body struct {
		MetamaskAddress string 	`json:"metamask_address"`
	}

	json.NewDecoder(r.Body).Decode(&body)
	nonce := pkg.GenerateNonce()
	nonceMap[body.MetamaskAddress] = nonce
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
        "nonce": nonce,
    })
}