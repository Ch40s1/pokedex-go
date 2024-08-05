package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val      []byte
	createAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:      val,
		createAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

// func (c *Cache) reap(interval time.Duration) {
// 	for k, v := range c.cache {
// 		if v.createAt.Before(t) {
// 			delete(c.cache, k)
// 		}
// 	}
// }
