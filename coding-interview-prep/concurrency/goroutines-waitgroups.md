# Concurrency: Goroutines & WaitGroups

Go's approach to concurrency is based on Communicating Sequential Processes (CSP), a formal language for describing patterns of interaction in concurrent systems. The two primary primitives for concurrency in Go are **goroutines** and **channels**.

## Goroutines

A **goroutine** is a lightweight thread of execution. They are functions or methods that run concurrently with other functions or methods.
-   **Lightweight:** Goroutines are multiplexed onto a smaller number of OS threads. They are much cheaper to create and manage than traditional threads.
-   **Managed by Go runtime:** The Go runtime handles scheduling, creation, and destruction of goroutines.
-   **Syntax:** To run a function as a goroutine, simply prepend the `go` keyword to the function call: `go myFunction()`.

## WaitGroups (`sync.WaitGroup`)

When you launch multiple goroutines, you often need a way to wait for them to complete their execution before proceeding with the rest of your program. The `sync.WaitGroup` type provides this synchronization mechanism.

A `WaitGroup` works as a counter:
-   `Add(delta int)`: Increments the counter by `delta`. You typically call this before launching a goroutine.
-   `Done()`: Decrements the counter by 1. You typically call this at the end of a goroutine (often using `defer`).
-   `Wait()`: Blocks until the counter becomes zero.

## Why use WaitGroups?
Without a `WaitGroup`, the `main` goroutine might exit before other goroutines have a chance to finish, leading to incomplete or incorrect results. `WaitGroup` ensures that the `main` goroutine waits for all specified goroutines to complete.

## Go Example: Basic Goroutines and WaitGroup Synchronization

See the example code in `goroutines-waitgroups.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run goroutines-waitgroups.go
```
