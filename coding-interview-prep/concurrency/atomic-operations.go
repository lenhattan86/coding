package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe" // Required for atomic.Pointer
)

// --- Atomic Counter Example ---

// Counter demonstrates an atomic counter.
func AtomicCounter() {
	var ops uint64 // Use uint64 for atomic operations
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1) // Atomically increment ops
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic Counter: Final ops count: %d\n", atomic.LoadUint64(&ops)) // Atomically load ops
}

// --- Atomic Pointer Swap Example ---

// Config represents a configuration struct.
type Config struct {
	Version string
	Settings map[string]string
}

// AtomicConfig demonstrates atomically swapping a pointer to a Config struct.
func AtomicConfig() {
	// Initial config
	cfg1 := &Config{
		Version: "v1.0",
		Settings: map[string]string{"timeout": "10s", "retries": "3"},
	}

	// atomic.Pointer[T] is a generic type for atomic operations on pointers to T.
	// It's safer than unsafe.Pointer for type-specific operations.
	var currentConfig atomic.Pointer[Config]
	currentConfig.Store(cfg1) // Store the initial config pointer

	// Reader goroutine
	go func() {
		for i := 0; i < 5; i++ {
			c := currentConfig.Load() // Atomically load the current config pointer
			fmt.Printf("Reader: Loaded config version: %s, settings: %v\n", c.Version, c.Settings)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Writer goroutine
	go func() {
		time.Sleep(200 * time.Millisecond) // Give reader a head start
		cfg2 := &Config{
			Version: "v2.0",
			Settings: map[string]string{"timeout": "20s", "retries": "5", "featureX": "enabled"},
		}
		currentConfig.Store(cfg2) // Atomically store the new config pointer
		fmt.Println("Writer: Updated config to v2.0")
	}()

	time.Sleep(1 * time.Second) // Let goroutines run
}

// --- CompareAndSwap Example (simplified) ---
func CompareAndSwapExample() {
	var value int32 = 0 // Must be a 32-bit or 64-bit integer for atomic ops

	// Try to change value from 0 to 1
	swapped := atomic.CompareAndSwapInt32(&value, 0, 1)
	fmt.Printf("CAS 1: Value: %d, Swapped: %t\n", value, swapped) // Expected: Value: 1, Swapped: true

	// Try to change value from 0 to 2 (will fail because value is now 1)
	swapped = atomic.CompareAndSwapInt32(&value, 0, 2)
	fmt.Printf("CAS 2: Value: %d, Swapped: %t\n", value, swapped) // Expected: Value: 1, Swapped: false

	// Try to change value from 1 to 2
	swapped = atomic.CompareAndSwapInt32(&value, 1, 2)
	fmt.Printf("CAS 3: Value: %d, Swapped: %t\n", value, swapped) // Expected: Value: 2, Swapped: true
}

func main() {
	fmt.Println("--- Atomic Operations Examples ---")

	fmt.Println("\n--- Atomic Counter ---")
	AtomicCounter()

	fmt.Println("\n--- Atomic Config (Pointer Swap) ---")
	AtomicConfig()

	fmt.Println("\n--- CompareAndSwap Example ---")
	CompareAndSwapExample()

	fmt.Println("--------------------")
}
