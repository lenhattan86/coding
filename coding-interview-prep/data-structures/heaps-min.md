# Heaps: Min-Heap

A heap is a specialized tree-based data structure that satisfies the heap property. In a **Min-Heap**, for any given node `C`, if `P` is a parent node of `C`, then the value of `P` is less than or equal to the value of `C`. The smallest element is always at the root.

Heaps are typically implemented using an array, where the tree structure is implicitly maintained.
-   For a node at index `i`:
    -   Its left child is at `2*i + 1`.
    -   Its right child is at `2*i + 2`.
    -   Its parent is at `(i - 1) / 2`.

## Characteristics:
-   It is a complete binary tree (all levels are completely filled except possibly the last level, and the last level has all its nodes as far left as possible).
-   Satisfies the heap property (min-heap or max-heap).
-   Efficiently finds the minimum (or maximum) element.

## Key Operations:
-   **Insert (Push):** Adds a new element to the heap while maintaining the heap property. This involves adding the element to the end and then "heapifying up" (swapping with parent if smaller) until the heap property is restored.
-   **Extract Min (Pop):** Removes and returns the minimum element (root) from the heap. This involves replacing the root with the last element, removing the last element, and then "heapifying down" (swapping with the smaller child if necessary) until the heap property is restored.
-   **Peek Min:** Returns the minimum element without removing it.
-   **Build Heap:** Converts an arbitrary array into a heap.

## Time Complexity:
-   **Insert:** O(log n)
-   **Extract Min:** O(log n)
-   **Peek Min:** O(1)
-   **Build Heap:** O(n)

## Use Cases:
-   Implementing Priority Queues.
-   Heap Sort algorithm.
-   Finding the k-th smallest/largest element.
-   Graph algorithms like Dijkstra's and Prim's (for efficient extraction of minimum edge/vertex).

## Go Example: Min-Heap Implementation

See the example code in `heaps-min.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run heaps-min.go
```
