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
	// var c eth.ContractFunctions
	// c.NewOrg("0x1A5b0307F532cd664f93D71786aa84b67964e635")
	// c.IsVerfiedOrg("0x1A5b0307F532cd664f93D71786aa84b67964e635")
	// t := c.CurrentTokenID()
	// c.TokenURI(t)
	r := router.RegisterRouter()
	// handlers.AllOrgs()
	//redis
	// rdb := redisdb.GetRedisInstance()
	// nonce := rdb.RedisGetNonce("0x5C67fAb678297594E35080831847084F8804E7Ae")
	
	// log.Println("Nonce: ", nonce)
	fmt.Println("Port :8080 is active....")
	http.ListenAndServe(":8080", r)
}
