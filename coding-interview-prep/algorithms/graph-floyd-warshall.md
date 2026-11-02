# Graph Algorithms: Floyd-Warshall Algorithm (All-Pairs Shortest Path)

The **Floyd-Warshall Algorithm** is a dynamic programming algorithm for finding shortest paths between all pairs of vertices in a weighted, directed graph. It can handle graphs with negative edge weights, but it cannot handle negative cycles (a cycle where the sum of edge weights is negative), as such a cycle would imply infinitely decreasing path lengths.

## Key Concepts:
-   **All-Pairs Shortest Path:** Finding the shortest path from every vertex to every other vertex in the graph.
-   **Dynamic Programming:** The algorithm builds up the solution by considering intermediate vertices.
-   **Adjacency Matrix:** Typically implemented using an adjacency matrix to store edge weights and path distances.

## How Floyd-Warshall Algorithm Works:
The algorithm works by iteratively improving an estimate of the shortest path between two vertices `i` and `j`. It considers all possible intermediate vertices `k` (from `0` to `V-1`, where `V` is the number of vertices).

Let `dist[i][j]` be the shortest distance from vertex `i` to vertex `j`.
The core idea is:
`dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])`

This means, the shortest path from `i` to `j` is either the direct path from `i` to `j`, or the path from `i` to `k` and then from `k` to `j`. The algorithm iterates through all possible `k` as intermediate vertices.

## Steps:
1.  **Initialization:** Create a distance matrix `dist[V][V]`.
    -   `dist[i][j]` is initialized to the weight of the edge from `i` to `j` if an edge exists.
    -   `dist[i][j]` is initialized to infinity if no direct edge exists.
    -   `dist[i][i]` is initialized to 0.
2.  **Iteration:** For each vertex `k` from `0` to `V-1`:
    -   For each vertex `i` from `0` to `V-1`:
        -   For each vertex `j` from `0` to `V-1`:
            -   `dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])`
3.  **Negative Cycle Detection:** After the algorithm completes, if `dist[i][i]` is negative for any `i`, then there is a negative cycle reachable from `i`.

## Time Complexity:
-   O(V^3) (where V is the number of vertices), due to three nested loops.

## Space Complexity:
-   O(V^2) for the distance matrix.

## Use Cases:
-   Finding all-pairs shortest paths in dense graphs.
-   Detecting negative cycles in a graph.
-   Transitive closure of a graph.
-   Used in various network routing protocols.

## Go Example: Floyd-Warshall Algorithm Implementation

See the example code in `graph-floyd-warshall.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run graph-floyd-warshall.go
```
