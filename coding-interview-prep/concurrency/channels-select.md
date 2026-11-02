# Concurrency: Select Statement

The `select` statement in Go is a powerful construct that allows a goroutine to wait on multiple communication operations (send or receive) simultaneously. It's similar to a `switch` statement but for channels.

## Key Concepts:
-   **Waiting on Multiple Channels:** `select` allows a goroutine to block until one of its cases is ready to proceed.
-   **Non-deterministic Choice:** If multiple cases are ready, `select` chooses one at random.
-   **Blocking:** If no cases are ready, the `select` statement blocks until one becomes ready.
-   **`default` Case:** If a `default` case is present, and no other case is ready, the `select` statement executes the `default` case immediately without blocking. This makes `select` non-blocking.
-   **`case <-ctx.Done():`:** A common pattern to handle cancellation or timeouts using the `context` package.

## Structure of `select` Statement:

```go
select {
case <-channel1:
    // Code to execute when channel1 is ready
case value := <-channel2:
    // Code to execute when channel2 is ready and a value is received
case channel3 <- valueToSend:
    // Code to execute when channel3 is ready to receive a value
case <-time.After(duration): // Timeout
    // Code to execute if no other case is ready within 'duration'
default:
    // Code to execute if no other case is ready immediately
}
```

## How it Works:
1.  **Evaluation:** All channel expressions in the `case` statements are evaluated.
2.  **Selection:**
    -   If one or more `case` statements are ready, one of them is chosen pseudo-randomly and its corresponding block of code is executed.
    -   If no `case` is ready:
        -   If a `default` case is present, it is executed.
        -   If no `default` case is present, the `select` statement blocks until one of the `case` statements becomes ready.

## Use Cases:
-   **Handling multiple incoming messages:** A server goroutine might listen on several request channels.
-   **Implementing timeouts:** Using `time.After` to limit how long an operation waits for a channel.
-   **Graceful shutdown/cancellation:** Listening for a cancellation signal from a `context.Done()` channel.
-   **Non-blocking channel operations:** Using `default` to try sending/receiving without blocking.
-   **Fan-in/Fan-out patterns:** Combining results from multiple goroutines or distributing tasks.

## Go Example: Select with Multiple Channels, Timeout, and Default

See the example code in `channels-select.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run channels-select.go
```
