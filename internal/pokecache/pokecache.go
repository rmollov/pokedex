package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cE       map[string]cacheEntry
	mX       sync.Mutex
	interval time.Duration
}

func NewCache(tt time.Duration) *Cache {

	items := make(map[string]cacheEntry)

	newC := &Cache{
		cE:       items,
		interval: tt,
	}

	go newC.reapLoop()
	return newC

}

func (c *Cache) Add(key string, val []byte) {

	c.mX.Lock()
	defer c.mX.Unlock()

	c.cE[key] = cacheEntry{createdAt: time.Now(), val: val}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mX.Lock()
	defer c.mX.Unlock()

	entry, ok := c.cE[key]

	if !ok {
		return nil, false
	}

	return entry.val, true

}

func (c *Cache) reapLoop() {

	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		cutoff := time.Now().Add(-c.interval)

		c.mX.Lock()

		for key, entry := range c.cE {

			if entry.createdAt.Before(cutoff) {
				delete(c.cE, key)
			}
		}
		c.mX.Unlock()
	}
}
