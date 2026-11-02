package main

import (
	"fmt"
	"time"
)

// producer sends numbers to a channel
func producer(id int, ch chan int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Duration(id*100) * time.Millisecond) // Simulate work
		ch <- id*10 + i
	}
	close(ch) // Close the channel when done
}

// consumer uses select to receive from multiple channels
func consumer(ch1, ch2 chan int) {
	for {
		select {
		case val1, ok := <-ch1:
			if !ok {
				ch1 = nil // Set channel to nil to stop receiving from it
				fmt.Println("Channel 1 closed.")
				break
			}
			fmt.Printf("Received from ch1: %d\n", val1)
		case val2, ok := <-ch2:
			if !ok {
				ch2 = nil // Set channel to nil to stop receiving from it
				fmt.Println("Channel 2 closed.")
				break
			}
			fmt.Printf("Received from ch2: %d\n", val2)
		case <-time.After(500 * time.Millisecond): // Timeout case
			fmt.Println("Consumer: No data received for 500ms. Timeout!")
			// If both channels are closed, exit the loop
			if ch1 == nil && ch2 == nil {
				fmt.Println("Consumer: All channels closed, exiting.")
				return
			}
		// default: // Uncommenting this makes select non-blocking
		// 	fmt.Println("Consumer: No channel ready, doing other work...")
		// 	Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("--- Select Statement Examples ---")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(1, ch1)
	go producer(2, ch2)

	consumer(ch1, ch2)

	fmt.Println("\n--- Select Statement with Timeout and Default (Non-blocking) ---")
	ch3 := make(chan string, 1)
	ch4 := make(chan string, 1)

	// Send to ch3 immediately
	ch3 <- "Hello"

	select {
	case msg1 := <-ch3:
		fmt.Printf("Received from ch3: %s\n", msg1)
	case msg2 := <-ch4:
		fmt.Printf("Received from ch4: %s\n", msg2)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout after 100ms (no data on ch4).")
default:
		fmt.Println("No channel ready immediately (ch3 was ready, but select picked it).")
	}

	// Example of default case when no channel is ready
	select {
	case msg1 := <-ch4:
		fmt.Printf("Received from ch4: %s\n", msg1)
default:
		fmt.Println("No channel ready immediately (ch4 is empty). Executing default.")
	}

	fmt.Println("--------------------")
}
