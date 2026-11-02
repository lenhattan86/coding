# Searching Algorithms

Searching algorithms are used to find an element or retrieve information from any data structure where the information is stored.

## Key Concepts:
-   **Target Element:** The specific value we are looking for.
-   **Data Structure:** The collection of elements being searched (e.g., array, linked list, tree).
-   **Time Complexity:** Measures how the running time of an algorithm grows with the input size (n).
-   **Space Complexity:** Measures how much extra memory an algorithm needs.

## Common Searching Algorithms:

### 1. Linear Search (Sequential Search)
-   **Concept:** The simplest searching algorithm. It sequentially checks each element of the list until a match is found or the whole list has been searched.
-   **Pre-requisites:** No specific order or structure required for the data.
-   **Time Complexity:**
    -   **Worst Case:** O(n) (target at the end or not present)
    -   **Average Case:** O(n)
    -   **Best Case:** O(1) (target at the beginning)
-   **Space Complexity:** O(1)

### 2. Binary Search
-   **Concept:** An efficient algorithm for finding an item from a *sorted* list of items. It works by repeatedly dividing in half the portion of the list that could contain the item, until you've narrowed down the possible locations to just one.
-   **Pre-requisites:** The data structure must be **sorted**.
-   **Time Complexity:**
    -   **Worst Case:** O(log n)
    -   **Average Case:** O(log n)
    -   **Best Case:** O(1) (target at the middle)
-   **Space Complexity:** O(1) (iterative) or O(log n) (recursive due to call stack)

## Go Example: Implementations of Searching Algorithms

See the example code in `searching.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run searching.go
```
