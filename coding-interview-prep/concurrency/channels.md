# Concurrency: Channels

Channels are the conduits through which you can send and receive values with a channel operator `<-`. They are a fundamental primitive for communication and synchronization between goroutines in Go.

## Key Concepts:
-   **Communication:** Channels allow goroutines to communicate safely by sending and receiving data.
-   **Synchronization:** Channels provide a way to synchronize the execution of goroutines. A send on a channel blocks until a receiver is ready, and a receive on a channel blocks until a sender is ready. This is known as "synchronous communication" or "rendezvous".

## Types of Channels:

1.  **Unbuffered Channels:**
    -   Created with `make(chan Type)`.
    -   A send operation on an unbuffered channel blocks until a receiver is ready to receive the value.
    -   A receive operation on an unbuffered channel blocks until a sender is ready to send a value.
    -   They provide a guarantee that an exchange of a value happens only when both sender and receiver are ready.

2.  **Buffered Channels:**
    -   Created with `make(chan Type, capacity)`. `capacity` specifies the number of elements the channel can hold.
    -   A send operation on a buffered channel blocks only when the buffer is full.
    -   A receive operation on a buffered channel blocks only when the buffer is empty.
    -   They allow goroutines to send and receive values asynchronously up to the buffer capacity.

## Channel Operations:
-   **Send:** `channel <- value`
-   **Receive:** `value := <-channel` or `<-channel` (to discard the value)
-   **Close:** `close(channel)`: Indicates that no more values will be sent on the channel. Receivers can check if a channel is closed.
-   **Range:** You can use a `for...range` loop to receive values from a channel until it is closed.

## Go Example: Unbuffered and Buffered Channels

See the example code in `channels.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run channels.go
```
