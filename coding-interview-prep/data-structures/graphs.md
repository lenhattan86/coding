# Graphs

A graph is a non-linear data structure consisting of a finite set of vertices (or nodes) and a set of edges that connect pairs of vertices. Graphs are used to model many real-world systems, such as social networks, road networks, and computer networks.

## Key Concepts:
-   **Vertex (Node):** A fundamental entity in a graph.
-   **Edge:** A connection between two vertices.
-   **Adjacency:** Two vertices are adjacent if they are connected by an edge.
-   **Path:** A sequence of vertices such that each consecutive pair is connected by an edge.
-   **Cycle:** A path that starts and ends at the same vertex.
-   **Directed Graph (Digraph):** Edges have a direction (e.g., A -> B means you can go from A to B, but not necessarily from B to A).
-   **Undirected Graph:** Edges have no direction (e.g., A -- B means you can go from A to B and B to A).
-   **Weighted Graph:** Edges have associated numerical values (weights or costs).
-   **Connected Graph:** In an undirected graph, there is a path between every pair of vertices.
-   **Strongly Connected Graph:** In a directed graph, there is a path from each vertex to every other vertex.
-   **Tree:** A connected acyclic undirected graph.

## Graph Representations:

1.  **Adjacency Matrix:**
    -   A 2D array where `matrix[i][j]` is 1 (or weight) if there is an edge from vertex `i` to vertex `j`, and 0 otherwise.
    -   **Pros:** Easy to implement, quick to check if an edge exists (O(1)).
    -   **Cons:** Space complexity is O(V^2) (V = number of vertices), even for sparse graphs.
2.  **Adjacency List:**
    -   An array of lists (or maps). The `i`-th element of the array contains a list of vertices adjacent to vertex `i`.
    -   **Pros:** Space complexity is O(V + E) (E = number of edges), efficient for sparse graphs. Easy to find all neighbors of a vertex.
    -   **Cons:** Checking if an edge exists takes O(degree of vertex) time.

## Graph Traversal Algorithms:

1.  **Breadth-First Search (BFS):**
    -   Explores all the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level.
    -   Uses a **Queue** data structure.
    -   Useful for finding the shortest path in an unweighted graph.
2.  **Depth-First Search (DFS):**
    -   Explores as far as possible along each branch before backtracking.
    -   Uses a **Stack** data structure (or recursion, which uses the call stack).
    -   Useful for topological sorting, finding connected components, and detecting cycles.

## Time Complexity:
-   **BFS/DFS (Adjacency List):** O(V + E)
-   **BFS/DFS (Adjacency Matrix):** O(V^2)

## Use Cases:
-   Social networks (friend connections).
-   GPS navigation (shortest path).
-   Network routing.
-   Dependency resolution (e.g., build systems).
-   Web crawling.

## Go Example: Graph Representation and Traversal

See the example code in `graphs.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run graphs.go
```
