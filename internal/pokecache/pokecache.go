package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex //locks the cache map, to prevent multible access (Mutex)
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {

	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()         //locks the cache map, to prevent multible access (Mutex)
	defer c.mux.Unlock() //unlocks the cache map after the Add is done
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()         //locks the cache map, to prevent multible access (Mutex)
	defer c.mux.Unlock() //unlocks the cache map after the Add is done
	cachE, ok := c.cache[key]
	return cachE.val, ok
}

func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()         //locks the cache map, to prevent multible access (Mutex)
	defer c.mux.Unlock() //unlocks the cache map after the Add is done
	timepast := time.Now().UTC().Add(-interval)
	for key, value := range c.cache {
		if value.createdAt.Before(timepast) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	//do not lock the infinity loop, otherwise it will be locked forever
	ticker := time.NewTicker(interval)
	for range ticker.C { //triggers every intervall the reap() e.g. every 5 sec
		c.reap(interval)
	}
}
