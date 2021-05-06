package bcache

import (
	"sync"
	"time"

	"github.com/darmiel/yaxc/internal/common"
	"github.com/muesli/termenv"
)

var prefix termenv.Style

func init() {
	p := common.Profile()
	prefix = termenv.String("CCHE").Foreground(p.Color("0")).Background(p.Color("#D290E4"))
}

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
