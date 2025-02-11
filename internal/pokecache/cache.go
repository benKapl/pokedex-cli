package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	c.cacheEntry[key] = entry
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.cacheEntry[key]
	return entry.val, exists
}

func (c *Cache) reapLoop() {
	c.mu.Lock()

	for key, entry := range c.cacheEntry {
		if time.Since(entry.createdAt) > c.interval {
			delete(c.cacheEntry, key)
		}
	}

	c.mu.Unlock()

}
