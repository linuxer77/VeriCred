package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vericred/pkg"
	"vericred/redisdb"
)

var nonceMap = map[string]string{}

func GetNonce(w http.ResponseWriter, r *http.Request) {
	var body struct {
		MetamaskAddress string 	`json:"metamask_address"`
	}

	fmt.Println("GetNonce function gets called.")
	json.NewDecoder(r.Body).Decode(&body)
	nonce := pkg.GenerateNonce()

	var c redisdb.Redis
	c.RedisSetNonce(body.MetamaskAddress, nonce)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
        "nonce": nonce,
    })
}