# Greedy Algorithms

A greedy algorithm is an algorithmic paradigm that follows the problem-solving heuristic of making the locally optimal choice at each stage with the hope of finding a global optimum. In many problems, a greedy strategy does not produce an optimal solution, but for some problems, it does.

## Key Characteristics:
-   **Greedy Choice Property:** A globally optimal solution can be arrived at by making a locally optimal (greedy) choice.
-   **Optimal Substructure:** An optimal solution to the problem contains optimal solutions to subproblems. (Similar to DP, but greedy makes a choice and doesn't look back).

## When to Use Greedy Algorithms:
Greedy algorithms are typically used for optimization problems. The main challenge is to prove that the greedy choice property holds for a given problem. If it does, a greedy algorithm can be much simpler and more efficient than dynamic programming.

## Difference from Dynamic Programming:
-   **Greedy:** Makes a choice that looks best at the moment and then solves the subproblem arising from that choice. The choice made by a greedy algorithm may depend on previous choices but never on future choices or all possible solutions. It commits to a choice.
-   **Dynamic Programming:** Explores all possible choices at each step and then picks the best one. It typically solves subproblems and then combines their solutions to get the solution for the main problem. It considers all possibilities.

## Common Greedy Problems:
-   Activity Selection Problem.
-   Fractional Knapsack Problem.
-   Huffman Coding.
-   Dijkstra's Algorithm (for shortest path).
-   Prim's and Kruskal's Algorithms (for Minimum Spanning Tree).
-   Coin Change Problem (sometimes, but not always optimal).

## Go Example: Activity Selection Problem

See the example code in `greedy-algorithms.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run greedy-algorithms.go
```
