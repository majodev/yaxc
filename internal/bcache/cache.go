package bcache

import (
	"log"
	"sync"
	"time"
)

type node struct {
	expires nodeExpiration
	value   interface{}
}

type Cache struct {
	values            sync.Map
	defaultExpiration time.Duration
	cleanerInterval   time.Duration
	f map[MapKey]string
}
type MapKey []byte

func NewCache(defaultExpiration, cleanerInterval time.Duration) *Cache {
	c := &Cache{
		defaultExpiration: defaultExpiration,
		cleanerInterval:   cleanerInterval,
	}
	if cleanerInterval != 0 {
		go c.janitorService()
	}
	return c
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	printDebugSet(key, value)
	c.values.Store([]byte(key), &node{
		expires: c.expiration(expiration),
		value:   value,
	})
}

func (c *Cache) Get(key string) (interface{}, bool) {
	log.Println("Cache; GET", key)
	var v interface{}
	var o bool
	if v, o = c.values.Load([]byte(key)); !o {
		log.Println("-> not in list")
		return nil, false
	}
	// cast
	var n *node
	if n, o = v.(*node); !o {
		log.Println("-> not a node")
		return nil, false
	}
	// expired
	if n.expires.IsExpired() {
		log.Println("-> expired")
		return nil, false
	}
	return n.value, true
}
