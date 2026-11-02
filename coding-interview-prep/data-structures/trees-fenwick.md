# Trees: Fenwick Tree (Binary Indexed Tree - BIT)

A **Fenwick Tree**, also known as a **Binary Indexed Tree (BIT)**, is a data structure that can efficiently update elements and calculate prefix sums in a table of numbers. It achieves this in logarithmic time, making it an alternative to Segment Trees for certain types of range queries.

## Characteristics:
-   **Efficient Prefix Sums:** Can calculate the sum of elements from index 0 to `i` in O(log n) time.
-   **Efficient Updates:** Can update an element at a specific index in O(log n) time.
-   **Space Complexity:** Requires O(n) space, similar to the input array.
-   **Simpler Implementation:** Generally simpler to implement than a Segment Tree for basic prefix sum/range sum operations.
-   **Limited Operations:** Primarily designed for prefix sums and single-point updates. More complex range operations (like range updates) are harder or impossible without modifications.

## Structure:
-   A Fenwick Tree is an array-based data structure.
-   Each index in the Fenwick Tree array `BIT[i]` stores the sum of a certain range of elements from the original array. The range represented by `BIT[i]` is determined by the least significant bit (LSB) of `i`.
-   The LSB of `i` (i.e., `i & -i`) gives the size of the range that `BIT[i]` is responsible for.

## Key Operations:
1.  **Update:** Adds a value `val` to the element at `idx` in the original array. This change needs to be propagated to all relevant nodes in the Fenwick Tree.
    -   Start at `idx + 1` (Fenwick Trees are typically 1-indexed).
    -   Add `val` to `BIT[idx]`.
    -   Move to the next index by adding `idx & -idx` to `idx`. Repeat until `idx` exceeds `n`.
2.  **Query (Prefix Sum):** Calculates the sum of elements from index 0 to `idx` in the original array.
    -   Start at `idx + 1`.
    -   Add `BIT[idx]` to the sum.
    -   Move to the previous index by subtracting `idx & -idx` from `idx`. Repeat until `idx` becomes 0.
3.  **Range Sum Query:** To find the sum of a range `[l, r]`, it's `query(r) - query(l-1)`.

## Time Complexity:
Let `n` be the size of the input array.
-   **Build (from scratch):** O(n log n) by repeatedly calling update, or O(n) with a specialized build function.
-   **Update:** O(log n)
-   **Query (Prefix Sum):** O(log n)

## Use Cases:
-   Calculating prefix sums and range sums.
-   Counting inversions in an array.
-   Dynamic updates and queries on frequency arrays.
-   Used in competitive programming for problems involving point updates and range queries.

## Go Example: Fenwick Tree for Prefix Sum Query

See the example code in `trees-fenwick.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run trees-fenwick.go
```
