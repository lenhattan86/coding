package main

import (
	"context"
	"fmt"
	"time"
)

// workerWithCancel simulates a task that can be cancelled
func workerWithCancel(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // Check if the context has been cancelled
			fmt.Printf("Worker %d: Cancelled. Error: %v\n", id, ctx.Err())
			return
		case <-time.After(500 * time.Millisecond): // Simulate doing some work
			fmt.Printf("Worker %d: Doing work...\n", id)
		}
	}
}

// workerWithTimeout simulates a task that might timeout
func workerWithTimeout(ctx context.Context, id int) {
	select {
	case <-time.After(2 * time.Second): // Simulate a long-running task
		fmt.Printf("Worker %d: Task completed successfully.\n", id)
	case <-ctx.Done(): // Check if the context has timed out or been cancelled
		fmt.Printf("Worker %d: Timed out/Cancelled. Error: %v\n", id, ctx.Err())
	}
}

// processRequest simulates processing a request with a value from context
func processRequest(ctx context.Context) {
	// Retrieve value from context
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}
	fmt.Printf("Processing request with ID: %s\n", requestID)

	// Simulate some work
	time.Sleep(100 * time.Millisecond)

	// Pass context to a sub-function
	subProcess(ctx)
}

func subProcess(ctx context.Context) {
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}
	fmt.Printf("Sub-process for request ID: %s\n", requestID)
}

func main() {
	fmt.Println("--- Context Package Examples ---")

	// --- Example 1: context.WithCancel ---
	fmt.Println("\n--- context.WithCancel ---")
	ctxCancel, cancel := context.WithCancel(context.Background())
	go workerWithCancel(ctxCancel, 1)
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Cancelling worker 1...")
	cancel() // Signal cancellation
	time.Sleep(500 * time.Millisecond) // Give time for worker to react

	// --- Example 2: context.WithTimeout ---
	fmt.Println("\n--- context.WithTimeout ---")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelTimeout() // Ensure resources are released
	go workerWithTimeout(ctxTimeout, 2)
	time.Sleep(1500 * time.Millisecond) // Wait longer than timeout
	fmt.Println("Main: Timeout example finished.")

	// --- Example 3: context.WithValue ---
	fmt.Println("\n--- context.WithValue ---")
	ctxValue := context.WithValue(context.Background(), "requestID", "req-12345")
	processRequest(ctxValue)
	fmt.Println("Main: Value example finished.")

	fmt.Println("--------------------")
}
