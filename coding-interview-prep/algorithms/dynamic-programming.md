# Dynamic Programming (DP)

Dynamic Programming is an algorithmic technique for solving a complex problem by breaking it down into simpler subproblems, solving each subproblem just once, and storing their solutions. The next time the same subproblem occurs, instead of recomputing its solution, we simply look up the already computed solution.

## Key Concepts:
-   **Optimal Substructure:** An optimal solution to the problem contains optimal solutions to subproblems.
-   **Overlapping Subproblems:** The same subproblems are encountered again and again.

## Approaches to Dynamic Programming:

1.  **Memoization (Top-Down):**
    -   Starts with the original problem and recursively breaks it down into subproblems.
    -   Stores the results of subproblems in a cache (e.g., map or array) as they are computed.
    -   If a subproblem's solution is already in the cache, it's retrieved directly.
    -   Often more intuitive to implement as it follows the natural recursive structure of the problem.

2.  **Tabulation (Bottom-Up):**
    -   Starts by solving the smallest subproblems first and then builds up solutions to larger subproblems.
    -   Typically uses an iterative approach (loops) to fill up a DP table (array or matrix).
    -   Avoids recursion overhead, potentially more efficient in terms of space and time for some problems.

## Steps to Solve a DP Problem:
1.  **Identify if it's a DP problem:** Look for optimal substructure and overlapping subproblems.
2.  **Define the state:** What does `DP[i]` or `DP[i][j]` represent?
3.  **Formulate the recurrence relation:** How can `DP[i]` be computed from previously computed states?
4.  **Identify the base cases:** What are the smallest subproblems and their solutions?
5.  **Determine the order of computation:** For tabulation, in what order should the DP table be filled?
6.  **Implement:** Choose memoization or tabulation.

## Time and Space Complexity:
-   Generally, DP solutions reduce exponential time complexities to polynomial time complexities (e.g., O(2^n) to O(n) or O(n^2)).
-   Space complexity is usually proportional to the size of the DP table.

## Use Cases:
-   Fibonacci sequence.
-   Longest Common Subsequence (LCS).
-   Knapsack problem.
-   Coin Change problem.
-   Matrix Chain Multiplication.
-   Edit Distance.
-   Many graph problems (e.g., shortest path on DAGs).

## Go Example: Fibonacci Sequence (Memoization & Tabulation) and Longest Common Subsequence

See the example code in `dynamic-programming.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run dynamic-programming.go
```
