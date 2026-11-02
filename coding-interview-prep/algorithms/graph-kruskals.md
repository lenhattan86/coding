# Graph Algorithms: Kruskal's Algorithm (Minimum Spanning Tree)

**Kruskal's Algorithm** is another greedy algorithm used to find a **Minimum Spanning Tree (MST)** for a weighted undirected graph. It differs from Prim's algorithm in its approach: instead of growing a single tree, Kruskal's algorithm builds the MST by adding edges one by one, always picking the cheapest available edge that connects two previously disconnected components.

## Key Concepts:
-   **Spanning Tree:** A subgraph that is a tree and connects all the vertices of the original graph.
-   **Minimum Spanning Tree (MST):** A spanning tree with the smallest possible sum of edge weights.
-   **Cut Property:** (Same as Prim's)
-   **Disjoint Set Union (DSU) / Union-Find Data Structure:** Crucial for efficiently detecting cycles. It keeps track of elements partitioned into a number of disjoint (non-overlapping) sets. It has two primary operations:
    -   **`Find(element)`:** Determines which set an element belongs to (by returning a representative of that set).
    -   **`Union(set1, set2)`:** Merges two sets into a single set.

## How Kruskal's Algorithm Works:
1.  **Initialization:** Create a list of all edges in the graph. Initialize a Disjoint Set Union (DSU) structure where each vertex is in its own set.
2.  **Sort Edges:** Sort all edges in non-decreasing order of their weights.
3.  **Iterate and Add:** Iterate through the sorted edges:
    -   For each edge `(u, v)` with weight `w`:
        -   If `u` and `v` are in different sets (i.e., adding this edge does not form a cycle), then add the edge to the MST, and merge the sets containing `u` and `v` using the `Union` operation.
4.  **Termination:** Stop when `V-1` edges have been added to the MST (where `V` is the number of vertices), or when all edges have been processed.

## Data Structures Used:
-   **List of Edges:** To store all edges and sort them.
-   **Disjoint Set Union (DSU):** To efficiently check for cycles and merge components.

## Time Complexity:
-   **Sorting Edges:** O(E log E) (where E is the number of edges). Since E can be at most V^2, this is also O(E log V).
-   **DSU Operations:** Each `Find` and `Union` operation takes nearly constant time on average (amortized O(α(V)), where α is the inverse Ackermann function, which grows extremely slowly).
-   **Overall:** Dominated by sorting, so O(E log E) or O(E log V).

## Go Example: Kruskal's Algorithm Implementation with DSU

See the example code in `graph-kruskals.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run graph-kruskals.go
```
