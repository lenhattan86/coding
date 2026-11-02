package main

import (
	"fmt"
	"sync"
	"time"
)

// worker simulates a task that takes some time
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes

	fmt.Printf("Worker %d starting...\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Printf("Worker %d finished.\n", id)
}

// anotherWorker demonstrates a goroutine without a WaitGroup (to show the problem)
func anotherWorker(id int) {
	fmt.Printf("Another Worker %d starting...\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Simulate work
	fmt.Printf("Another Worker %d finished.\n", id)
}

func main() {
	fmt.Println("--- Goroutines and WaitGroups ---")

	// --- Example 1: Goroutines without WaitGroup (demonstrates race condition/early exit) ---
	fmt.Println("\n--- Goroutines without WaitGroup (might not see all output) ---")
	go anotherWorker(1)
	go anotherWorker(2)
	// The main goroutine might exit before anotherWorker(1) or anotherWorker(2) finish.
	// To see their output, we'd need to add a sleep here, which is not ideal.
	// time.Sleep(3 * time.Second) // Uncomment to see output, but this is a hack.
	fmt.Println("Main goroutine finished (without waiting for 'anotherWorker').")

	// --- Example 2: Goroutines with WaitGroup ---
	fmt.Println("\n--- Goroutines with WaitGroup (ensures all workers complete) ---")
	var wg sync.WaitGroup // Declare a WaitGroup

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the counter for each goroutine launched
		go worker(i, &wg) // Pass the WaitGroup by reference
	}

	fmt.Println("Main goroutine waiting for workers to finish...")
	wg.Wait() // Block until the WaitGroup counter is zero
	fmt.Println("All workers finished. Main goroutine continuing.")

	fmt.Println("--------------------")
}
