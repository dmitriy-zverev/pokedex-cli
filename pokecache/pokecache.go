package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data:       make(map[string]cacheEntry),
		data_mutex: &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
