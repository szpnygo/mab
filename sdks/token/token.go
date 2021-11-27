package token

import (
	"context"
	"time"
)

type TokenStorageHelper interface {
	// SetToken 设置token
	SetToken(ctx context.Context, key, token string, expiration time.Duration) error
	// GetToken 获取token
	GetToken(ctx context.Context, key string) (string, error)
}
