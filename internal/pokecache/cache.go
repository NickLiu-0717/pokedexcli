package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	cachemap map[string]cacheEntry
	mu       *sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	cache := cacheEntry{
		createAt: time.Now().UTC(),
		val:      val,
	}
	c.cachemap[key] = cache
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ca, ok := c.cachemap[key]
	if ok {
		return ca.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cachemap {
			now := time.Now().UTC()
			if entry.createAt.Before(now.Add(-interval)) {
				delete(c.cachemap, key)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cachemap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go cache.reaploop(interval)
	return cache
}
