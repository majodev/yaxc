package bcache

import (
	"log"
	"time"
)

func (c *Cache) janitorService() {
	if c.cleanerInterval == 0 {
		return
	}
	for {
		time.Sleep(c.cleanerInterval)
		printDebugJanitorStart()
		c.janitor()
	}
}

func (c *Cache) janitor() {
	c.values.Range(func(key, value interface{}) bool {
		log.Println("* K:", key, "; V:", value)
		if value == nil {
			c.values.Delete(key)
			log.Println("-> * Deleted. (nil):", key)
			return true
		}
		// cast
		if node, ok := value.(*node); ! ok {
			c.values.Delete(key)
			log.Println("-> * Deleted. (cast):", key)
			return true
		} else {
			// expired?
			if node.expires.IsExpired() {
				c.values.Delete(key)
				log.Println("-> * Deleted. (expired):", key)
				return true
			}
		}
		return true
	})
}
