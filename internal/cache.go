package internal

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]CacheEntry
	mu    *sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(d time.Duration) Cache {
	cache := Cache{
		Entry: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}

	go reapLoop(d, &cache)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	// Lock accessing of cache when writing
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entry[key] = CacheEntry{
		Val:       val,
		CreatedAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if m, ok := c.Entry[key]; ok {
		return m.Val, true
	}

	return []byte{}, false
}

// Is called once a new cache is created.
// Keeps checking based on the interval and remove any entries
// older than the interval.
func reapLoop(d time.Duration, c *Cache) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for {
		t := <-ticker.C

		for k, v := range c.Entry {
			if t.Sub(v.CreatedAt) > d {
				delete(c.Entry, k)
			}
		}
	}
}
