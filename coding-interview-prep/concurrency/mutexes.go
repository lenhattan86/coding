package main

import (
	"fmt"
	"sync"
	"time"
)

// --- Mutex Example ---

// SafeCounter is a counter that can be safely used by multiple goroutines.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// --- RWMutex Example ---

// SafeData is a data structure protected by an RWMutex.
type SafeData struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewSafeData creates a new SafeData instance.
func NewSafeData() *SafeData {
	return &SafeData{
		data: make(map[string]string),
	}
}

// Read reads a value from the data map.
func (sd *SafeData) Read(key string) (string, bool) {
	sd.mu.RLock() // Acquire a read lock
	defer sd.mu.RUnlock() // Release the read lock
	val, ok := sd.data[key]
	return val, ok
}

// Write writes a value to the data map.
func (sd *SafeData) Write(key, value string) {
	sd.mu.Lock() // Acquire a write lock
	defer sd.mu.Unlock() // Release the write lock
	sd.data[key] = value
}

func main() {
	fmt.Println("--- Mutex Example ---")
	c := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter Value (Mutex): %d\n", c.Value()) // Expected: 1000
	fmt.Println("--------------------")

	fmt.Println("\n--- RWMutex Example ---")
	sd := NewSafeData()

	// Writers
	go func() {
		sd.Write("name", "Alice")
		time.Sleep(50 * time.Millisecond)
		sd.Write("city", "New York")
	}()

	// Readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id*10) * time.Millisecond) // Stagger reads
			name, ok := sd.Read("name")
			if ok {
				fmt.Printf("Reader %d: Read name = %s\n", id, name)
			} else {
				fmt.Printf("Reader %d: Name not found\n", id)
			}
			city, ok := sd.Read("city")
			if ok {
				fmt.Printf("Reader %d: Read city = %s\n", id, city)
			} else {
				fmt.Printf("Reader %d: City not found\n", id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("--------------------")
}
