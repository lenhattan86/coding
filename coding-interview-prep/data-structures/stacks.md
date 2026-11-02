# Stacks

A stack is a linear data structure that follows a particular order in which operations are performed. The order is LIFO (Last In, First Out) or FILO (First In, Last Out). This means the last element added to the stack is the first one to be removed.

## Key Operations:
-   **Push:** Adds an element to the top of the stack.
-   **Pop:** Removes the top element from the stack.
-   **Peek (or Top):** Returns the top element of the stack without removing it.
-   **IsEmpty:** Checks if the stack is empty.
-   **Size:** Returns the number of elements in the stack.

## Characteristics:
-   LIFO/FILO principle.
-   Can be implemented using arrays (dynamic arrays/slices in Go) or linked lists.

## Time Complexity:
-   **Push:** O(1)
-   **Pop:** O(1)
-   **Peek:** O(1)
-   **IsEmpty:** O(1)

## Use Cases:
-   Function call stack in programming languages.
-   Undo/Redo functionality in applications.
-   Browser history (back button).
-   Expression evaluation (e.g., infix to postfix conversion).
-   Backtracking algorithms.

## Go Example: Stack Implementation

See the example code in `stacks.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run stacks.go
```
