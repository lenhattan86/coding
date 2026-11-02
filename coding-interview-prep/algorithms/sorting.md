# Sorting Algorithms

Sorting algorithms are algorithms that put elements of a list in a certain order. The most frequently used orders are numerical order and lexicographical order.

## Key Concepts:
-   **In-place vs. Out-of-place:** An in-place algorithm transforms input using a small, constant amount of extra space. An out-of-place algorithm requires extra space proportional to the input size.
-   **Stable vs. Unstable:** A sorting algorithm is stable if it preserves the relative order of equal elements.
-   **Comparison Sort vs. Non-Comparison Sort:** Comparison sorts rely on comparing elements. Non-comparison sorts (like Radix, Counting) do not use comparisons and can achieve better time complexities for specific data distributions.
-   **Time Complexity:** Measures how the running time of an algorithm grows with the input size (n).
-   **Space Complexity:** Measures how much extra memory an algorithm needs.

## Common Sorting Algorithms:

### 1. Bubble Sort
-   **Concept:** Repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
-   **Time Complexity:** O(n^2) (Worst, Average, Best)
-   **Space Complexity:** O(1)
-   **Stable:** Yes

### 2. Selection Sort
-   **Concept:** Divides the input list into two parts: a sorted sublist and an unsorted sublist. It repeatedly finds the minimum element from the unsorted sublist and puts it at the end of the sorted sublist.
-   **Time Complexity:** O(n^2) (Worst, Average, Best)
-   **Space Complexity:** O(1)
-   **Stable:** No

### 3. Insertion Sort
-   **Concept:** Builds the final sorted array (or list) one item at a time. It iterates through the input elements and removes one element at a time, finds the place within the sorted list, and inserts it there.
-   **Time Complexity:** O(n^2) (Worst, Average), O(n) (Best)
-   **Space Complexity:** O(1)
-   **Stable:** Yes

### 4. Merge Sort
-   **Concept:** A divide-and-conquer algorithm. It divides the unsorted list into n sublists, each containing one element, and then repeatedly merges sublists to produce new sorted sublists until there is only one sorted list remaining.
-   **Time Complexity:** O(n log n) (Worst, Average, Best)
-   **Space Complexity:** O(n)
-   **Stable:** Yes

### 5. Quick Sort
-   **Concept:** A divide-and-conquer algorithm. It picks an element as a pivot and partitions the array around the picked pivot.
-   **Time Complexity:** O(n log n) (Average, Best), O(n^2) (Worst)
-   **Space Complexity:** O(log n) (Average), O(n) (Worst)
-   **Stable:** No

### 6. Heap Sort
-   **Concept:** A comparison-based sorting technique based on the Binary Heap data structure. It's similar to selection sort where we first find the maximum element and place it at the end. We repeat the same process for the remaining elements.
-   **Time Complexity:** O(n log n) (Worst, Average, Best)
-   **Space Complexity:** O(1)
-   **Stable:** No

## Go Example: Implementations of Sorting Algorithms

See the example code in `sorting.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run sorting.go
```
