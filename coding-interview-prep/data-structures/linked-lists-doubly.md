# Linked Lists: Doubly Linked List

A **Doubly Linked List** is a variation of the linked list in which navigation is possible in both directions, either forward or backward. This is achieved by having each node contain two pointers: one pointing to the next node in the sequence (like a singly linked list) and another pointing to the previous node.

## Node Structure:
Each node in a doubly linked list typically consists of:
-   **Data:** The value stored in the node.
-   **`next` pointer:** A reference to the next node in the list.
-   **`prev` pointer:** A reference to the previous node in the list.

## Characteristics:
-   **Bidirectional Traversal:** Can be traversed from head to tail and from tail to head.
-   **Efficient Deletion:** Deleting a node is more efficient than in a singly linked list because you have direct access to the previous node, eliminating the need to traverse from the head to find the predecessor.
-   **Increased Memory Overhead:** Each node requires an extra pointer (`prev`), leading to more memory consumption compared to a singly linked list.

## Key Operations:
-   **Insertion:**
    -   At the beginning.
    -   At the end.
    -   After a given node.
    -   Before a given node.
-   **Deletion:**
    -   From the beginning.
    -   From the end.
    -   A specific node.
-   **Traversal:**
    -   Forward (from head to tail).
    -   Backward (from tail to head).

## Time Complexity:
-   **Access (by index):** O(n)
-   **Search:** O(n)
-   **Insertion (at beginning/end/given node):** O(1) (assuming you have a pointer to the insertion point)
-   **Deletion (at beginning/end/given node):** O(1) (assuming you have a pointer to the node to delete)

## Use Cases:
-   Implementing LRU (Least Recently Used) caches.
-   Navigation systems (e.g., browser history: back/forward buttons).
-   Representing decks of cards.
-   Implementing other data structures like hash tables with separate chaining where efficient deletion is needed.

## Go Example: Doubly Linked List Implementation

See the example code in `linked-lists-doubly.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run linked-lists-doubly.go
```
