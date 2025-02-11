package pokecache

import (
	"sync"
	"time"
)

type (
	cacheEntry struct {
		createdAt time.Time
		val       []byte
	}
	Cache struct {
		cache map[string]cacheEntry
		mu    *sync.Mutex
	}
)

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cache, exists := c.cache[key]
	return cache.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if time.Since(v.createdAt) > interval {
			delete(c.cache, k)
		}
	}

}
