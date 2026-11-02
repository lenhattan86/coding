# Trees: Binary Search Trees (BST)

A tree is a non-linear data structure that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.

A **Binary Search Tree (BST)**, also known as an ordered or sorted binary tree, is a special type of binary tree where for each node:
-   The value of all nodes in its left subtree is less than the node's value.
-   The value of all nodes in its right subtree is greater than the node's value.
-   Both the left and right subtrees are also binary search trees.

## Characteristics:
-   Nodes are ordered, which allows for efficient searching.
-   The left child is always less than the parent, and the right child is always greater.
-   Can become unbalanced, leading to worst-case time complexities.

## Key Operations:
-   **Insert:** Adds a new node while maintaining the BST property.
-   **Search:** Finds a node with a specific value.
-   **Delete:** Removes a node while maintaining the BST property. This is the most complex operation, especially for nodes with two children.
-   **Traversal:** Visiting all nodes in a specific order:
    -   **In-order:** Left -> Root -> Right (produces sorted output)
    -   **Pre-order:** Root -> Left -> Right
    -   **Post-order:** Left -> Right -> Root

## Time Complexity:
-   **Average Case (Balanced Tree):**
    -   **Insert:** O(log n)
    -   **Search:** O(log n)
    -   **Delete:** O(log n)
-   **Worst Case (Unbalanced/Skewed Tree):**
    -   **Insert:** O(n)
    -   **Search:** O(n)
    -   **Delete:** O(n)

## Use Cases:
-   Implementing dynamic sets or maps.
-   Used as a building block for more complex data structures like AVL trees and Red-Black trees (self-balancing BSTs).
-   Database indexing.

## Go Example: Binary Search Tree Implementation

See the example code in `trees-bst.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run trees-bst.go
```
