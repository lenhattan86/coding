# Queues

A queue is a linear data structure that follows a particular order in which operations are performed. The order is FIFO (First In, First Out). This means the first element added to the queue is the first one to be removed.

## Key Operations:
-   **Enqueue (or Offer):** Adds an element to the rear (end) of the queue.
-   **Dequeue (or Poll):** Removes the front (beginning) element from the queue.
-   **Peek (or Front):** Returns the front element of the queue without removing it.
-   **IsEmpty:** Checks if the queue is empty.
-   **Size:** Returns the number of elements in the queue.

## Characteristics:
-   FIFO principle.
-   Can be implemented using arrays (dynamic arrays/slices in Go) or linked lists. Linked list implementation is generally preferred for dynamic resizing and efficient enqueue/dequeue operations.

## Time Complexity:
-   **Enqueue:** O(1)
-   **Dequeue:** O(1)
-   **Peek:** O(1)
-   **IsEmpty:** O(1)

## Use Cases:
-   CPU scheduling.
-   Printer queues.
-   Call center phone systems.
-   Breadth-First Search (BFS) algorithm.
-   Buffering in streaming data.

## Go Example: Queue Implementation

See the example code in `queues.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run queues.go
```
