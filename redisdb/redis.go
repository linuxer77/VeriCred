package redisdb

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

var (
	instance *Redis
	once     sync.Once
)

func GetRedisInstance() *Redis {
	once.Do(func() {
		instance = &Redis{
			Client: redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
				Protocol: 1,
			}),
		}
	})
	return instance
}

func (r *Redis) RedisSetNonce(key, value string) {
	ctx := context.Background()
	fmt.Println("Key: ", key)
	fmt.Println("Value: ", value)
	
	err := r.Client.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully set the key value pair in redis.")
	val, err := r.Client.Get(ctx, key).Result()
	
	if err != nil {
		log.Println("Error: ", err)
	}
	log.Println("Val: ", val)
}

func (r *Redis) RedisGetNonce(key string) string {
	ctx := context.Background()
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("Error: ", err)
	}
	
	return val
}