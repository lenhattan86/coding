# Recursion & Backtracking

## Recursion

Recursion is a programming technique where a function calls itself directly or indirectly to solve a problem. A recursive function solves a problem by breaking it down into smaller, identical subproblems until it reaches a base case, which can be solved directly without further recursion.

### Key Concepts:
-   **Base Case:** The condition that stops the recursion. Without a base case, a recursive function would call itself indefinitely, leading to a stack overflow.
-   **Recursive Step:** The part of the function where it calls itself with a modified input, moving closer to the base case.

### Characteristics:
-   Often leads to more elegant and concise code for problems that have a recursive structure (e.g., tree traversals, factorial, Fibonacci).
-   Can be less efficient than iterative solutions due to function call overhead.
-   Can lead to stack overflow errors if the recursion depth is too large.

### Use Cases:
-   Tree and graph traversals (DFS).
-   Divide and Conquer algorithms (Merge Sort, Quick Sort).
-   Mathematical functions (factorial, Fibonacci).
-   Parsing and language processing.

## Backtracking

Backtracking is an algorithmic technique for solving problems, typically constraint satisfaction problems, by trying to build a solution incrementally. It explores all possible solutions by systematically trying to extend a partial solution. If a partial solution cannot be extended to a complete solution, or if it violates any constraints, the algorithm "backtracks" (undoes its last step) and tries a different path.

### Key Concepts:
-   **State Space Tree:** The conceptual tree representing all possible partial and complete solutions. Backtracking explores this tree.
-   **Constraints:** Rules that must be satisfied by a valid solution.
-   **Pruning:** Eliminating branches of the state space tree that cannot lead to a valid solution, improving efficiency.

### Characteristics:
-   Often implemented using recursion.
-   Explores all potential candidates for a solution.
-   When a candidate fails, it backtracks to the previous state.

### Use Cases:
-   N-Queens problem.
-   Sudoku Solver.
-   Knight's Tour problem.
-   Combinatorial optimization problems (e.g., finding all permutations, combinations).
-   Solving mazes.

## Go Example: Factorial (Recursion) and N-Queens (Backtracking)

See the example code in `recursion-backtracking.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run recursion-backtracking.go
```
