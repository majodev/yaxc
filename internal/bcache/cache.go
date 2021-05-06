package bcache

import (
	"sync"
	"time"
)

type node struct {
	value interface{}
}

type Cache struct {
	mu     sync.Mutex
	values map[string]*node
}

func NewCache() *Cache {
	c := &Cache{
		values: make(map[string]*node),
	}
	return c
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.Lock()

	c.values[key] = &node{
		value: value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	if v, o := c.values[key]; o && v != nil {
		c.mu.Unlock()
		return v.value, true
	}
	c.mu.Unlock()
	return nil, false
}
