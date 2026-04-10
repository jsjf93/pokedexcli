package pokecache

import (
	"errors"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type SafeCache struct {
	mu    *sync.RWMutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) SafeCache {
	cache := SafeCache{
		mu:    &sync.RWMutex{},
		cache: make(map[string]cacheEntry),
	}
	go cache.cleanUpCache(interval)
	return cache
}

func (c *SafeCache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.cache[key]; ok {
		return errors.New("cache entry with key: " + key + " already exists")
	}

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	return nil
}

func (c *SafeCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *SafeCache) cleanUpCache(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.cache {
			if time.Since(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}
