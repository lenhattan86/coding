package main

import (
	"fmt"
	"sync"
	"time"
)

// DataItem represents a piece of data we want to cache.
type DataItem struct {
	ID    string
	Value string
	// Add other fields as needed
}

// MockDatabase simulates a slow database.
type MockDatabase struct {
	data map[string]DataItem
	mu   sync.RWMutex
}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{
		data: make(map[string]DataItem),
	}
}

// Fetch simulates fetching data from the database.
func (db *MockDatabase) Fetch(id string) (DataItem, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	fmt.Printf("DB: Fetching item %s...\n", id)
	time.Sleep(200 * time.Millisecond) // Simulate network latency and DB query time
	item, ok := db.data[id]
	return item, ok
}

// Store simulates storing data in the database.
func (db *MockDatabase) Store(item DataItem) {
	db.mu.Lock()
	defer db.mu.Unlock()
	fmt.Printf("DB: Storing item %s...\n", item.ID)
	time.Sleep(100 * time.Millisecond) // Simulate write time
	db.data[item.ID] = item
}

// CacheEntry holds cached data and its expiration time.
type CacheEntry struct {
	Item       DataItem
	Expiration time.Time
}

// SimpleCache simulates a cache with LRU eviction and TTL.
type SimpleCache struct {
	cache    map[string]CacheEntry
	mu       sync.RWMutex
	ttl      time.Duration // Time-To-Live for cache entries
	capacity int
	lruList  *list.List // For LRU eviction
}

func NewSimpleCache(capacity int, ttl time.Duration) *SimpleCache {
	return &SimpleCache{
		cache:    make(map[string]CacheEntry),
		ttl:      ttl,
		capacity: capacity,
		lruList:  list.New(),
	}
}

// Get implements the read part of Cache-Aside.
func (c *SimpleCache) Get(id string, db *MockDatabase) (DataItem, bool) {
	c.mu.RLock()
	entry, found := c.cache[id]
	c.mu.RUnlock()

	if found && time.Now().Before(entry.Expiration) {
		fmt.Printf("Cache: Hit for %s\n", id)
		// Update LRU position (move to front/back)
		c.mu.Lock()
		for e := c.lruList.Front(); e != nil; e = e.Next() {
			if e.Value.(string) == id {
				c.lruList.MoveToFront(e)
				break
			}
		}
		c.mu.Unlock()
		return entry.Item, true
	}

	if found {
		fmt.Printf("Cache: Miss for %s (stale)\n", id)
		// Invalidate stale entry
		c.mu.Lock()
		delete(c.cache, id)
		for e := c.lruList.Front(); e != nil; e = e.Next() {
			if e.Value.(string) == id {
				c.lruList.Remove(e)
				break
			}
		}
		c.mu.Unlock()
	} else {
		fmt.Printf("Cache: Miss for %s\n", id)
	}

	// Fetch from DB
	item, ok := db.Fetch(id)
	if ok {
		c.mu.Lock()
		// Evict if capacity reached
		if c.lruList.Len() >= c.capacity {
			oldestID := c.lruList.Back().Value.(string)
			delete(c.cache, oldestID)
			c.lruList.Remove(c.lruList.Back())
			fmt.Printf("Cache: Evicted %s (LRU)\n", oldestID)
		}
		// Add to cache
		c.cache[id] = CacheEntry{Item: item, Expiration: time.Now().Add(c.ttl)}
		c.lruList.PushFront(id)
		c.mu.Unlock()
	}
	return item, ok
}

// Invalidate implements the write part of Cache-Aside.
func (c *SimpleCache) Invalidate(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, found := c.cache[id]; found {
		delete(c.cache, id)
		for e := c.lruList.Front(); e != nil; e = e.Next() {
			if e.Value.(string) == id {
				c.lruList.Remove(e)
				break
			}
		}
		fmt.Printf("Cache: Invalidated %s\n", id)
	}
}

func main() {
	fmt.Println("--- Conceptual Caching Strategies (Cache-Aside) ---")

	db := NewMockDatabase()
	cache := NewSimpleCache(2, 500*time.Millisecond) // Capacity 2, TTL 500ms

	// Initial data in DB
	db.Store(DataItem{ID: "user1", Value: "Alice"})
	db.Store(DataItem{ID: "user2", Value: "Bob"})
	db.Store(DataItem{ID: "user3", Value: "Charlie"})

	fmt.Println("\n--- First Reads (Cache Misses) ---")
	cache.Get("user1", db) // Miss, fetch from DB, add to cache
	cache.Get("user2", db) // Miss, fetch from DB, add to cache

	fmt.Println("\n--- Second Reads (Cache Hits) ---")
	cache.Get("user1", db) // Hit
	cache.Get("user2", db) // Hit

	fmt.Println("\n--- Read causing Eviction ---")
	cache.Get("user3", db) // Miss, fetch from DB, add to cache, user1 evicted (LRU)

	fmt.Println("\n--- Read after Eviction ---")
	cache.Get("user1", db) // Miss again (was evicted)

	fmt.Println("\n--- Stale Data (TTL Expiration) ---")
	time.Sleep(600 * time.Millisecond) // Wait for TTL to expire
	cache.Get("user2", db)             // Miss (stale), fetch from DB

	fmt.Println("\n--- Write Operation (Invalidation) ---")
	updatedUser2 := DataItem{ID: "user2", Value: "Bobby"}
	db.Store(updatedUser2) // Write to DB
	cache.Invalidate("user2") // Invalidate cache entry

	cache.Get("user2", db) // Miss (invalidated), fetch from DB

	fmt.Println("\n--- Caching Strategies Example Finished ---")
}
