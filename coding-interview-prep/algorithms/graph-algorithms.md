# Graph Algorithms

Graph algorithms are a set of instructions that traverse or analyze graphs to solve specific problems. Graphs are powerful data structures used to model relationships between entities.

## Key Concepts Revisited:
-   **Vertex (Node):** An entity in the graph.
-   **Edge:** A connection between two vertices.
-   **Weighted Graph:** Edges have associated costs.
-   **Directed Graph:** Edges have a direction.
-   **Path:** A sequence of vertices connected by edges.
-   **Cycle:** A path that starts and ends at the same vertex.

## Common Graph Algorithms:

### 1. Dijkstra's Algorithm (Shortest Path)
-   **Purpose:** Finds the shortest paths from a single source vertex to all other vertices in a graph with non-negative edge weights.
-   **Approach:** A greedy algorithm that uses a priority queue to always extract the vertex with the smallest known distance.
-   **Time Complexity:**
    -   With a min-priority queue implemented using a binary heap: O(E log V) or O(E + V log V)
    -   With a Fibonacci heap: O(E + V log V)
-   **Limitations:** Does not work with negative edge weights.

### 2. Prim's Algorithm (Minimum Spanning Tree - MST)
-   **Purpose:** Finds a Minimum Spanning Tree for a weighted undirected graph. An MST is a subset of the edges of a connected, edge-weighted undirected graph that connects all the vertices together, without any cycles and with the minimum possible total edge weight.
-   **Approach:** A greedy algorithm that starts from an arbitrary vertex and grows the MST by adding the cheapest edge that connects a vertex in the MST to a vertex outside the MST.
-   **Time Complexity:**
    -   With adjacency matrix: O(V^2)
    -   With adjacency list and binary heap: O(E log V) or O(E + V log V)

### 3. Kruskal's Algorithm (Minimum Spanning Tree - MST)
-   **Purpose:** Also finds a Minimum Spanning Tree for a weighted undirected graph.
-   **Approach:** A greedy algorithm that sorts all the edges in non-decreasing order of their weights. It then adds edges to the MST one by one, if adding the edge does not form a cycle. A Disjoint Set Union (DSU) data structure is typically used to efficiently detect cycles.
-   **Time Complexity:** O(E log E) or O(E log V) (dominated by sorting edges)

### 4. Floyd-Warshall Algorithm (All-Pairs Shortest Path)
-   **Purpose:** Finds shortest paths between all pairs of vertices in a weighted directed graph. Can handle negative edge weights, but not negative cycles.
-   **Approach:** Dynamic programming. It considers all possible intermediate vertices between every pair of source and destination vertices.
-   **Time Complexity:** O(V^3)

### 5. Topological Sort
-   **Purpose:** For a Directed Acyclic Graph (DAG), it produces a linear ordering of its vertices such that for every directed edge `u -> v`, vertex `u` comes before `v` in the ordering.
-   **Approach:** Can be implemented using DFS or Kahn's algorithm (using BFS and in-degrees).
-   **Time Complexity:** O(V + E)

## Go Example: Dijkstra's Algorithm

See the example code in `graph-algorithms.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run graph-algorithms.go
```
