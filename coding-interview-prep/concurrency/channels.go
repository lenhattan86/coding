package main

import (
	"fmt"
	"time"
)

// --- Unbuffered Channel Example ---
func unbufferedSender(ch chan int) {
	fmt.Println("Unbuffered Sender: Sending 10")
	ch <- 10 // This will block until a receiver is ready
	fmt.Println("Unbuffered Sender: Sent 10")
}

func unbufferedReceiver(ch chan int) {
	fmt.Println("Unbuffered Receiver: Waiting to receive...")
	val := <-ch // This will block until a sender sends a value
	fmt.Printf("Unbuffered Receiver: Received %d\n", val)
}

// --- Buffered Channel Example ---
func bufferedSender(ch chan int) {
	fmt.Println("Buffered Sender: Sending 1")
	ch <- 1 // This will not block immediately if buffer has space
	fmt.Println("Buffered Sender: Sending 2")
	ch <- 2
	fmt.Println("Buffered Sender: Sending 3 (will block if buffer size is 2)")
	ch <- 3 // This might block if capacity is 2 and no receiver yet
	fmt.Println("Buffered Sender: Sent all values")
	close(ch) // Close the channel after sending all values
}

func bufferedReceiver(ch chan int) {
	fmt.Println("Buffered Receiver: Waiting to receive...")
	for val := range ch { // Loop until channel is closed and empty
		fmt.Printf("Buffered Receiver: Received %d\n", val)
		time.Sleep(50 * time.Millisecond) // Simulate some work
	}
	fmt.Println("Buffered Receiver: Channel closed and empty.")
}

func main() {
	fmt.Println("--- Channels ---")

	// --- Unbuffered Channel Demo ---
	fmt.Println("\n--- Unbuffered Channel Demo ---")
	unbufferedCh := make(chan int)
	go unbufferedSender(unbufferedCh)
	go unbufferedReceiver(unbufferedCh)
	time.Sleep(100 * time.Millisecond) // Give goroutines time to execute
	fmt.Println("Unbuffered channel demo finished.")

	// --- Buffered Channel Demo ---
	fmt.Println("\n--- Buffered Channel Demo (Capacity 2) ---")
	bufferedCh := make(chan int, 2) // Create a buffered channel with capacity 2
	go bufferedSender(bufferedCh)
	go bufferedReceiver(bufferedCh)
	time.Sleep(200 * time.Millisecond) // Give goroutines time to execute
	fmt.Println("Buffered channel demo finished.")

	fmt.Println("--------------------")
}
