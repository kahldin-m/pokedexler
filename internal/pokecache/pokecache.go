package pokecache

import (
	"sync"
	"time"
)

// Cache
type Cache struct {
	mu        sync.Mutex
	entries   map[string]cacheEntry
	duration  time.Duration
}

// cacheEntry
type cacheEntry struct {
	createdAt    time.Time
	val          []byte
}

// Constructor function for new Cache objects
func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:   make(map[string]cacheEntry),
		duration:  interval,
	}
	go c.reapLoop()
	return c
}

// Add a new entry to the cache
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{createdAt: time.Now(), val: value}
}

// Return an entry at key from stored cache
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entries[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

// Reap method is called when the cache is created by NewCache
func (c *Cache) reapLoop() {
	// make a ticker that fires every c.duration
	ticker := time.NewTicker(c.duration)
	// every time if fires, loop through every entry in the map
	for range ticker.C {
		// for each entry, compute how old it is by subracting createdAt from Now
		c.mu.Lock()
		for key, entry := range c.entries {
			age := time.Now().Sub(entry.createdAt)
			// if it's older than the allowed duration, remove it from the map by "key"
			if age > c.duration {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
