# Concurrency: Mutexes & RWMutexes

While channels are the idiomatic way to communicate and synchronize in Go, sometimes you need to protect shared memory directly. This is where mutexes come in.

## Mutex (`sync.Mutex`)

A **Mutex** (Mutual Exclusion) is a synchronization primitive that provides exclusive access to a shared resource. It ensures that only one goroutine can access a critical section of code at any given time, preventing race conditions.

-   **`Lock()`:** Acquires a lock. If the lock is already held by another goroutine, the calling goroutine blocks until the lock is available.
-   **`Unlock()`:** Releases a lock. It's crucial to call `Unlock()` after `Lock()` to allow other goroutines to acquire the lock. `defer mutex.Unlock()` is a common pattern to ensure the lock is released even if the function panics.

## RWMutex (`sync.RWMutex`)

A **RWMutex** (Read-Write Mutex) is a more specialized type of mutex that allows for multiple readers or a single writer. It's useful when you have a shared resource that is read frequently but written to infrequently.

-   **`Lock()`:** Acquires a write lock. This blocks all other readers and writers.
-   **`Unlock()`:** Releases a write lock.
-   **`RLock()`:** Acquires a read lock. Multiple goroutines can hold a read lock concurrently. However, if a write lock is held, `RLock()` blocks until the write lock is released. If a write lock is requested, `RLock()` will block until all current read locks are released.
-   **`RUnlock()`:** Releases a read lock.

## When to use which:
-   **`sync.Mutex`:** Use when you need to protect a resource that is modified frequently or when you have a simple critical section that needs exclusive access.
-   **`sync.RWMutex`:** Use when you have a resource that is read much more often than it is written. This can improve performance by allowing concurrent reads.

## Go Example: Mutex and RWMutex for Shared Counter

See the example code in `mutexes.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run mutexes.go
```
