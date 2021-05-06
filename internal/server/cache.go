package server

import (
	"errors"
	"sync"
)

type Cache struct {
	mu     sync.Mutex
	values map[string]*node
}

type node struct {
	value interface{}
}

func NewCache() *Cache {
	c := &Cache{
		values: make(map[string]*node),
	}
	return c
}

func (b *Cache) Get(key string) (res string, err error) {
	return b.get("val::" + key)
}

func (b *Cache) GetHash(key string) (res string, err error) {
	return b.get("hash::" + key)
}

func (b *Cache) Set(key, value string) error {
	b.setValue("val::"+key, value)
	return nil
}

func (b *Cache) SetHash(key, value string) error {
	b.setValue("hash::"+key, value)
	return nil
}

func (b *Cache) get(key string) (res string, err error) {
	if v, ok := b.getValue(key); ok {
		if r, o := v.(string); o {
			res = r
		} else {
			err = errors.New("cast error")
		}
	}
	return
}

func (c *Cache) setValue(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[key] = &node{
		value: value,
	}
}

func (c *Cache) getValue(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, o := c.values[key]; o && v != nil {
		return v.value, true
	}
	return nil, false
}
