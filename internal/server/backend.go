package server

import (
	"time"
)

type Backend interface {
	GetValue(key string) (string, error)
	GetHash(key string) (string, error)
	Set(key, value, hash string, ttl time.Duration) error
}
