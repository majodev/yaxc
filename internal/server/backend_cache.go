package server

import (
	"github.com/darmiel/yaxc/internal/bcache"
	"time"
)

type CacheBackend struct {
	cache   *bcache.Cache
	errCast error
	errNotFound error
}

type valueHashTuple struct {
	value, hash []byte
}

func (v *valueHashTuple) val(b []byte) string {
	if b == nil {
		return ""
	}
	return string(b)
}
func (v *valueHashTuple) Value() string {
	return v.val(v.value)
}
func (v *valueHashTuple) Hash() string {
	return v.val(v.hash)
}

func (b *CacheBackend) GetValue(key string) (res string, err error) {
	log.Debug("Requested value:", key)
	var t *valueHashTuple
	if t, err = b.getTuple(key); err != nil {
		return
	}
	return t.Value(), nil
}

func (b *CacheBackend) GetHash(key string) (res string, err error) {
	log.Debug("Requested hash:", key)
	var t *valueHashTuple
	if t, err = b.getTuple(key); err != nil {
		return
	}
	return t.Hash(), nil
}

func (b *CacheBackend) Set(key, value, hash string, ttl time.Duration) error {
	t := &valueHashTuple{
		value: []byte(value),
		hash:  []byte(hash),
	}
	b.cache.Set(key, t, ttl)
	return nil
}

func (b *CacheBackend) getTuple(key string) (*valueHashTuple, error) {
	var v interface{}
	var o bool
	// find in cache
	if v, o = b.cache.Get(key); !o {
		return nil, b.errNotFound
	}
	// cast
	var t *valueHashTuple
	if t, o = v.(*valueHashTuple); !o {
		return nil, b.errCast
	}
	return t, nil
}
