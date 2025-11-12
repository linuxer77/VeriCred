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
				Addr:     "redis-18234.c61.us-east-1-3.ec2.redns.redis-cloud.com:18234",
				Username: "default",
				Password: "2I1L2vRC7nTfGu04Gr6x7wLc2Sh3sZpR",
				DB:       0,
			}),
		}
	})

	ctx := context.Background()
	_, err := instance.Client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Redis connection failed: %v", err)
	}
	log.Println("Successfully connected to redis db")
	return instance
}

func (r *Redis) RedisSetNonce(key, value string) error {
	ctx := context.Background()

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
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			log.Printf("Redis key not found: %s", key)
			return "", fmt.Errorf("key not found: %s", key)
		}
		log.Printf("Failed to get Redis key %s: %v", key, err)
		return "", fmt.Errorf("failed to get redis key: %w", err)
	}

	return val, nil
}

