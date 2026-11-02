# Concurrency: Atomic Operations (`sync/atomic`)

While `sync.Mutex` and `sync.RWMutex` are excellent for protecting complex data structures and critical sections, sometimes you only need to protect a single variable (like an integer counter or a boolean flag) from race conditions. For such simple cases, using a mutex can introduce unnecessary overhead.

The `sync/atomic` package in Go provides low-level atomic memory primitives that guarantee operations on certain types of variables are performed as a single, indivisible unit, without interference from other goroutines. These operations are typically implemented using special CPU instructions, making them very efficient.

## Key Concepts:
-   **Atomicity:** An operation is atomic if it appears to occur instantaneously and indivisibly from the perspective of other goroutines. It either completes entirely or not at all, and no intermediate state is visible.
-   **Race Condition:** Occurs when multiple goroutines access shared memory concurrently, and at least one of them is a write operation, leading to unpredictable results. Atomic operations help prevent these for simple types.

## Functions in `sync/atomic`:

The `sync/atomic` package provides functions for various types, including `int32`, `int64`, `uint32`, `uint64`, `uintptr`, and `unsafe.Pointer`.

1.  **`AddInt32`, `AddInt64`, `AddUint32`, `AddUint64`:**
    -   Atomically adds a value to a variable and returns the new value.
    -   `atomic.AddInt64(&counter, 1)`

2.  **`LoadInt32`, `LoadInt64`, `LoadUint32`, `LoadUint64`, `LoadPointer`:**
    -   Atomically loads (reads) the value of a variable. Ensures that the read value is the most recently written value.
    -   `value := atomic.LoadInt64(&counter)`

3.  **`StoreInt32`, `StoreInt64`, `StoreUint32`, `StoreUint64`, `StorePointer`:**
    -   Atomically stores (writes) a new value to a variable. Ensures that the write is fully completed before any other goroutine can observe it.
    -   `atomic.StoreInt64(&counter, 0)`

4.  **`CompareAndSwapInt32`, `CompareAndSwapInt64`, `CompareAndSwapUint32`, `CompareAndSwapUint64`, `CompareAndSwapPointer`:**
    -   **CAS (Compare-And-Swap):** Atomically compares the value of a variable with an "old" value. If they are equal, it swaps the variable's value with a "new" value. It returns `true` if the swap occurred, `false` otherwise.
    -   This is a fundamental building block for many lock-free algorithms.
    -   `atomic.CompareAndSwapInt64(&counter, expected, new)`

## When to use Atomic Operations:
-   When you need to perform simple, single-variable operations (increment, decrement, read, write, compare-and-swap) on primitive types.
-   When performance is critical, as atomic operations are generally faster than mutexes for these specific tasks.
-   As building blocks for more complex lock-free data structures.

## Limitations:
-   Only works for specific primitive types and `unsafe.Pointer`.
-   Cannot protect complex data structures (e.g., maps, slices, structs) from race conditions; for those, mutexes or channels are still necessary.

## Go Example: Atomic Counter and Atomic Pointer Swap

See the example code in `atomic-operations.go`.
```go
// To run this code:
// 1. Navigate to the concurrency directory: cd coding-interview-prep/concurrency
// 2. Run the file: go run atomic-operations.go
```
