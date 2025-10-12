package pokecache

import (
	"sync"
	"time"
)

//Caching of previous results so we don't query API repeatedly
// for same result set.

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Notice mutex to protect the Entries map.
type Cache struct {
	Entries    map[string]cacheEntry
	CacheMutex sync.Mutex
}

// Pointer receiver as otherwise we will copy the struct and have a new mutex
// i.e copylock. But we would still have the same map (reference!) - so multiple
// different callers of Add would have their own separate mutexes but the same map reference
// which would allow multiple callers to simultaneously read/write to the map which is UB
// Add an entry to the Entries map.
func (c *Cache) Add(key string, val []byte) {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	c.Entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	entry, exists := c.Entries[key]
	if exists {
		return entry.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	//Wait for a tick from the channel - should be a tick every interval
	// Once we stop blocks - we can loop again.
	for {
		// I am assuming we will block here until tick received and removed.
		<-ticker.C

		//Not using the tick time - interval - in case we were blocked.
		// We want to collect anything that is more than interval old.
		cutoffTime := time.Now().Add(-interval)

		//Acquire lock on cache mutex - will block here until lock acquirable
		c.CacheMutex.Lock()

		//Loop through map, delete any k:v where val.createdAt is after cutoffTime
		for key, val := range c.Entries {
			if val.createdAt.Before(cutoffTime) {
				delete(c.Entries, key)
			}
		}
		//Drop lock
		c.CacheMutex.Unlock()
	}

}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		//need to initialise a map - rest fine as zeroinit
		Entries: make(map[string]cacheEntry),
	}
	//Calls looping method that deletes any entries older than
	// the interval every interval tick using time.Ticker
	go cache.reapLoop(interval)
	return cache
}
