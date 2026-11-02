# Linked Lists

A linked list is a linear data structure where elements are not stored at contiguous memory locations. Instead, each element (called a node) is an object that contains a data field and a reference (or pointer) to the next node in the sequence.

## Types of Linked Lists:

1.  **Singly Linked List:** Each node points to the next node in the sequence. The last node points to `nil`.
2.  **Doubly Linked List:** Each node has two pointers: one to the next node and one to the previous node. This allows traversal in both directions.
3.  **Circular Linked List:** The last node points back to the first node, forming a circle. This can be singly or doubly linked.

## Characteristics:
-   **Dynamic Size:** Linked lists can grow or shrink in size during runtime.
-   **Non-contiguous Memory:** Elements are not stored in adjacent memory locations.
-   **Efficient Insertions/Deletions:** Adding or removing elements is generally faster than arrays (O(1) once the position is found), as it only requires updating pointers.
-   **Inefficient Random Access:** Accessing an element by index requires traversing the list from the beginning (O(n)).

## Time Complexity:
-   **Access (by index):** O(n)
-   **Search:** O(n)
-   **Insertion (at beginning/end/given node):** O(1) (assuming you have a pointer to the insertion point)
-   **Deletion (at beginning/end/given node):** O(1) (assuming you have a pointer to the node to delete and its predecessor for singly linked lists)

## Use Cases:
-   Implementing stacks and queues.
-   Image viewers (previous and next image).
-   Web browser history (back and forward buttons).
-   Polynomial representation.

## Go Example: Singly Linked List

See the example code in `linked-lists.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run linked-lists.go
```
