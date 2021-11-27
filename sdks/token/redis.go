package token

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/szpnygo/mab/internal/errorx"
)

type RedisTokenStorage struct {
	*redis.Client
}

func NewRedisTokenStorage(redis *redis.Client) *RedisTokenStorage {
	return &RedisTokenStorage{
		Client: redis,
	}
}

func (s *RedisTokenStorage) SetToken(ctx context.Context, key, token string, expiration time.Duration) error {
	return s.Set(ctx, key, token, expiration).Err()
}

func (s *RedisTokenStorage) GetToken(ctx context.Context, key string) (string, error) {
	result, err := s.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errorx.New("token is empty")
	}

	return result, nil
}
