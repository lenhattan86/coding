# Graph Algorithms: Prim's Algorithm (Minimum Spanning Tree)

**Prim's Algorithm** is a greedy algorithm that finds a **Minimum Spanning Tree (MST)** for a weighted undirected graph. An MST is a subset of the edges of a connected, edge-weighted undirected graph that connects all the vertices together, without any cycles and with the minimum possible total edge weight.

## Key Concepts:
-   **Spanning Tree:** A subgraph that is a tree and connects all the vertices of the original graph.
-   **Minimum Spanning Tree (MST):** A spanning tree with the smallest possible sum of edge weights.
-   **Cut Property:** For any cut (a partition of the vertices into two disjoint sets), if an edge crosses the cut and has strictly smaller weight than any other edge crossing the cut, then this edge must be part of all MSTs. Prim's algorithm implicitly uses this property.

## How Prim's Algorithm Works:
1.  **Initialization:** Start with an arbitrary vertex as the initial MST.
2.  **Iteration:** In each step, find the edge with the minimum weight that connects a vertex in the current MST to a vertex outside the MST.
3.  **Expansion:** Add this minimum-weight edge and the new vertex to the MST.
4.  **Termination:** Repeat until all vertices are included in the MST.

## Data Structures Used:
-   **Adjacency List:** To represent the graph.
-   **Min-Priority Queue:** To efficiently find the minimum-weight edge connecting to the MST. The priority queue stores edges (or vertices with their minimum connecting edge weight) and allows quick extraction of the minimum.

## Time Complexity:
-   **Adjacency Matrix Implementation:** O(V^2) (where V is the number of vertices). This is efficient for dense graphs.
-   **Adjacency List + Binary Heap (Min-Priority Queue):** O(E log V) or O(E + V log V) (where E is the number of edges). This is more efficient for sparse graphs.

## Go Example: Prim's Algorithm Implementation

See the example code in `graph-prims.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run graph-prims.go
```
