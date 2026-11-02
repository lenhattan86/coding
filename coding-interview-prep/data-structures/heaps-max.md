# Heaps: Max-Heap

A heap is a specialized tree-based data structure that satisfies the heap property. In a **Max-Heap**, for any given node `C`, if `P` is a parent node of `C`, then the value of `P` is greater than or equal to the value of `C`. The largest element is always at the root.

Heaps are typically implemented using an array, where the tree structure is implicitly maintained.
-   For a node at index `i`:
    -   Its left child is at `2*i + 1`.
    -   Its right child is at `2*i + 2`.
    -   Its parent is at `(i - 1) / 2`.

## Characteristics:
-   It is a complete binary tree (all levels are completely filled except possibly the last level, and the last level has all its nodes as far left as possible).
-   Satisfies the heap property (min-heap or max-heap).
-   Efficiently finds the maximum (or minimum) element.

## Key Operations:
-   **Insert (Push):** Adds a new element to the heap while maintaining the heap property. This involves adding the element to the end and then "heapifying up" (swapping with parent if larger) until the heap property is restored.
-   **Extract Max (Pop):** Removes and returns the maximum element (root) from the heap. This involves replacing the root with the last element, removing the last element, and then "heapifying down" (swapping with the larger child if necessary) until the heap property is restored.
-   **Peek Max:** Returns the maximum element without removing it.
-   **Build Heap:** Converts an arbitrary array into a heap.

## Time Complexity:
-   **Insert:** O(log n)
-   **Extract Max:** O(log n)
-   **Peek Max:** O(1)
-   **Build Heap:** O(n)

## Use Cases:
-   Implementing Priority Queues (where highest priority means largest value).
-   Heap Sort algorithm (as seen in `sorting.go`).
-   Finding the k-th largest element.

## Go Example: Max-Heap Implementation

See the example code in `heaps-max.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run heaps-max.go
```
