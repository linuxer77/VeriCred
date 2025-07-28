package redisdb

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
}

func (r *Redis) RedisSetNonce(key, value string) {
	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 1,
	})

	ctx := context.Background()

	err := client.Set(ctx, key, value, time.Minute * 5).Err()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully set the key value pair in redis.")
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("Val: ", val)
}

func (r *Redis) RedisGetNonce(key string) string {
	var client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 1,
	})
	
	ctx := context.Background()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error: ", err)
	}
	
	return val
}