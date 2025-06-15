package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
    ctx = context.Background()
    rdb *redis.Client
)

func ConnectRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        panic(fmt.Sprintf("REDIS | Failed to connect: %v", err))
    }
    fmt.Println("REDIS | Connected successfully.")
}

func SetValue(key string, value string, ttlSeconds int) error {
    return rdb.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func GetValue(key string) (string, error) {
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("REDIS | key '%s' not found", key)
    }
    return val, err
}

func DeleteKey(key string) error {
    return rdb.Del(ctx, key).Err()
}

func IncrementValue(key string) (int64, error) {
    return rdb.Incr(ctx, key).Result()
}
