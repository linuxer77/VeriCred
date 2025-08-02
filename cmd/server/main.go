package main

import (
	"fmt"
	"log"
	"net/http"
	"vericred/internal/db"
	"vericred/internal/router"
	"vericred/redisdb"
)

func main() {
	db.Init()	
	r := router.RegisterRouter()

	//redis
	rdb := redisdb.GetRedisInstance()
	nonce := rdb.RedisGetNonce("0x5C67fAb678297594E35080831847084F8804E7Ae")
	
	log.Println("Nonce: ", nonce)
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
