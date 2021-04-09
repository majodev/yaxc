package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisBackend struct {
	ctx       context.Context
	client    *redis.Client
	prefixVal string
	prefixHsh string
}

func (b *RedisBackend) GetValue(key string) (res string, err error) {
	return b.get(b.prefixVal, key)
}

func (b *RedisBackend) GetHash(key string) (res string, err error) {
	return b.get(b.prefixHsh, key)
}

func (b *RedisBackend) Set(key, value, hash string, ttl time.Duration) (err error) {
	if err = b.set(b.prefixVal, key, value, ttl); err != nil {
		return
	}
	err = b.set(b.prefixHsh, key, hash, ttl)
	return
}

///

func (b *RedisBackend) get(prefix, key string) (res string, err error) {
	cmd := b.client.Get(b.ctx, prefix+key)
	if err := cmd.Err(); err != nil && err != redis.Nil {
		return "", err
	}
	res, _ = cmd.Result()
	return
}

func (b *RedisBackend) set(prefix, key, value string, ttl time.Duration) (err error) {
	cmd := b.client.Set(b.ctx, prefix+key, value, ttl)
	err = cmd.Err()
	return
}
