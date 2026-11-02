package main

import (
	"fmt"
	"sync"
	"time"
)

// --- Stage 1: Generator ---
// Generates a sequence of numbers and sends them to an output channel.
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// --- Stage 2: Squarer ---
// Reads numbers from an input channel, squares them, and sends to an output channel.
func squarer(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// --- Stage 3: Multiplier ---
// Reads numbers from an input channel, multiplies them by a factor, and sends to an output channel.
func multiplier(done <-chan struct{}, in <-chan int, factor int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * factor:
			case <-done:
				return
			}
		}
	}()
	return out
}

// --- Stage 4: Printer (Sink) ---
// Reads numbers from an input channel and prints them.
func printer(done <-chan struct{}, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		select {
		case <-done:
			return
		default:
			fmt.Printf("Result: %d\n", n)
		}
	}
}

func main() {
	fmt.Println("--- Concurrency Pipeline Example ---")

	done := make(chan struct{}) // Channel to signal cancellation
	defer close(done)           // Ensure done is closed when main exits

	// Stage 1: Generate numbers
	nums := []int{1, 2, 3, 4, 5}
	gen := generator(done, nums...)

	// Stage 2: Square the numbers
	sq := squarer(done, gen)

	// Stage 3: Multiply by a factor
	mult := multiplier(done, sq, 10)

	// Stage 4: Print results
	var wg sync.WaitGroup
	wg.Add(1)
	printer(done, mult, &wg)

	wg.Wait() // Wait for the printer to finish

	fmt.Println("--- Pipeline Example Finished ---")
}
