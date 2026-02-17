package redisdb

import (
	"context"
	"crypto/tls"
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
				Addr:     "redis-cqtv-chah-625658.leapcell.cloud:6379",
				Username: "default",
				Password: "Ae00000i0bCV9NoGMY4YyTLaWbKckBrUHeVsQUCEqtFJrLd/kDZKMxmzm13oNzh+Ixopp+a",
				DB:       0,
				TLSConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
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
