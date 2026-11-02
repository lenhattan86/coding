# Trees: Segment Tree

A **Segment Tree** is a tree data structure used for storing information about intervals or segments. It allows querying for information about a range (e.g., sum, minimum, maximum) in logarithmic time, and updating an element or a range in logarithmic time.

## Characteristics:
-   **Efficient Range Queries:** Can answer range queries (sum, min, max, etc.) in O(log n) time.
-   **Efficient Updates:** Can update a single element or a range in O(log n) time.
-   **Space Complexity:** Requires O(n) space, typically 2n or 4n depending on implementation details.

## Structure:
-   It's a binary tree where each node represents an interval.
-   The root node represents the entire array/range.
-   Each leaf node represents a single element of the input array.
-   Each internal node represents the union of its children's intervals.
-   The value stored in an internal node is typically the result of an operation (sum, min, max) over its corresponding interval.

## Key Operations:
1.  **Build:** Constructs the segment tree from an input array. This is typically done recursively.
    -   If the interval is a single element, the leaf node stores that element.
    -   Otherwise, it recursively builds left and right children for the two halves of the interval and stores the combined result.
2.  **Update:** Modifies the value of an element in the original array and propagates the change up the tree to maintain consistency.
3.  **Query:** Retrieves information (e.g., sum, min) for a given range. It traverses the tree, combining results from relevant nodes.

## Time Complexity:
Let `n` be the size of the input array.
-   **Build:** O(n)
-   **Update:** O(log n)
-   **Query:** O(log n)

## Use Cases:
-   Range Minimum/Maximum Query (RMQ).
-   Range Sum Query (RSQ).
-   Finding the number of elements in a range that satisfy a certain condition.
-   Used in competitive programming for problems involving intervals.

## Go Example: Segment Tree for Range Sum Query

See the example code in `trees-segment.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run trees-segment.go
```
