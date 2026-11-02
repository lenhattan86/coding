# Concurrency: Context Package

The `context` package in Go provides a way to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines. It's an essential tool for managing the lifecycle of operations, especially in complex, concurrent applications like web servers or microservices.

## Key Concepts:
-   **Cancellation:** A `Context` can be canceled, signaling to all goroutines listening to that context that they should stop their work.
-   **Deadlines:** A `Context` can have a deadline, after which it is automatically canceled. This is useful for operations that should not take longer than a certain amount of time.
-   **Values:** A `Context` can carry request-scoped values (e.g., user ID, trace ID) that are safe for concurrent use.

## `context.Context` Interface:
The `context.Context` interface defines four methods:
-   `Deadline() (deadline time.Time, ok bool)`: Returns the time when the context will be canceled, if any.
-   `Done() <-chan struct{}`: Returns a channel that is closed when the context is canceled or times out.
-   `Err() error`: Returns an error that indicates why the context was canceled.
-   `Value(key interface{}) interface{}`: Returns the value associated with `key`, or `nil` if no value is associated.

## Creating Contexts:
-   **`context.Background()`:** The root context for all contexts. It's never canceled, has no deadline, and carries no values. Used for the top-level scope of a program.
-   **`context.TODO()`:** Used when you are unsure which context to use or when the context is not yet available. It's a placeholder.

## Derived Contexts:
You derive new contexts from existing ones using functions like:
-   **`context.WithCancel(parent Context)`:** Returns a new `Context` and a `CancelFunc`. Calling the `CancelFunc` cancels the derived context and all its children.
-   **`context.WithDeadline(parent Context, d time.Time)`:** Returns a new `Context` and a `CancelFunc`. The derived context is canceled when the deadline `d` passes or when the parent context is canceled, whichever happens first.
-   **`context.WithTimeout(parent Context, timeout time.Duration)`:** Similar to `WithDeadline`, but takes a `time.Duration` for convenience.
-   **`context.WithValue(parent Context, key, val interface{})`:** Returns a new `Context` that carries the key-value pair.

## Go Example: Context with Cancellation, Timeout, and Value

See the example code in `context-package.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run context-package.go
```
