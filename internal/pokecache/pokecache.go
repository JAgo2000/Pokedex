package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cachE, ok := c.cache[key]
	return cachE.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	timepast := time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(timepast) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C { //triggers every intervall the reap() e.g. every 5 sec
		c.reap(interval)
	}
}
