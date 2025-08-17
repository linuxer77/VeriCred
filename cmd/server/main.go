package main

import (
	"fmt"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/logging"
	"vericred/internal/router"
)

func main() {
	db.Init()	
	logging.Init()

	r := router.RegisterRouter()
	// handlers.AllOrgs()
	//redis
	// rdb := redisdb.GetRedisInstance()
	// nonce := rdb.RedisGetNonce("0x5C67fAb678297594E35080831847084F8804E7Ae")
	
	// log.Println("Nonce: ", nonce)
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
