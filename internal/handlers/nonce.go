package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"vericred/pkg"
	"vericred/redisdb"
)

func GetNonce(w http.ResponseWriter, r *http.Request) {
	var body struct {
		MetamaskAddress string 	`json:"metamask_address"`
	}
   
	// bodyBytes, _ := io.ReadAll(r.Body)
	// fmt.Println("Raw JSON body:", string(bodyBytes))
	
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println("Some error occurred when decoding: ", err)
	}
	nonce := pkg.GenerateNonce()

	rdb := redisdb.GetRedisInstance()
	rdb.RedisSetNonce(body.MetamaskAddress, nonce)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
        "nonce": nonce,
    })
}