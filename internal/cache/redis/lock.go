package redis

import (
	"context"
	"time"
)

func AcquireLock(ctx context.Context, key string, expiration time.Duration) bool {
	success, err := RedisClient.SetNX(ctx, key, "locked", expiration).Result()
	if err != nil {
		return false
	}
	return success
}

func ReleaseLock(ctx context.Context, key string) {
	RedisClient.Del(ctx, key)
}
