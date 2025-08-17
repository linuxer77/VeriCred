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

	ctx := context.Background()
	_, err := instance.Client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
	} else {
		log.Println("Redis connected successfully")
	}
	return instance
}

func (r *Redis) RedisSetNonce(key, value string) error {
	ctx := context.Background()
	fmt.Println("Key: ", key)
	fmt.Println("Value: ", value)
	
	err := r.Client.Set(ctx, key, value, time.Minute*5).Err()
	if err != nil {
        log.Printf("Failed to set Redis key %s: %v", key, err)
        return fmt.Errorf("failed to set redis key: %w", err)
	}

    log.Printf("Successfully set Redis key: %s", key)
	return nil
	// val, err := r.Client.Get(ctx, key).Result()
	
	// if err != nil {
	// 	log.Println("Error: ", err)
	// }
	// log.Println("Val: ", val)
}

func (r *Redis) RedisGetNonce(key string) (string, error) {
	ctx := context.Background()
    log.Printf("Getting Redis key: %s", key)
	val, err := r.Client.Get(ctx, key).Result()

	if err != nil {
        if err == redis.Nil {
            log.Printf("Redis key not found: %s", key)
            return "", fmt.Errorf("key not found: %s", key)
        }
        log.Printf("Failed to get Redis key %s: %v", key, err)
        return "", fmt.Errorf("failed to get redis key: %w", err)
    }
    
    log.Printf("Successfully retrieved Redis key: %s = %s", key, val)
    return val, nil
}