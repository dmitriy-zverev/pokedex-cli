package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data       map[string]cacheEntry
	data_mutex *sync.Mutex
}

func (c *Cache) Add(key string, value []byte) {
	c.data_mutex.Lock()
	defer c.data_mutex.Unlock()
	c.data[key] = cacheEntry{createdAt: time.Now().UTC(), val: value}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.data_mutex.Lock()
	defer c.data_mutex.Unlock()
	cacheItem, ok := c.data[key]
	return cacheItem.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.data_mutex.Lock()
	defer c.data_mutex.Unlock()
	for key, value := range c.data {
		if value.createdAt.Before(now.Add(-interval)) {
			delete(c.data, key)
		}
	}
}
