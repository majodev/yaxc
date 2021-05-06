package bcache

import (
	"sync"
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

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[key] = &node{
		value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, o := c.values[key]; o && v != nil {
		return v.value, true
	}
	return nil, false
}
