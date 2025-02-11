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
		cacheEntry map[string]cacheEntry
		interval   time.Duration
		mu         *sync.Mutex
	}
)

func NewCache(interval time.Duration) Cache {
	return Cache{
		interval: interval,
		mu:       &sync.Mutex{},
	}
}
