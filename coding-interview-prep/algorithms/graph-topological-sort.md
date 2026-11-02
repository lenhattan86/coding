# Graph Algorithms: Topological Sort

**Topological Sort** (or Topological Ordering) is a linear ordering of vertices in a **Directed Acyclic Graph (DAG)** such that for every directed edge `u -> v`, vertex `u` comes before `v` in the ordering. A topological sort is only possible for DAGs; if a graph contains a cycle, a topological sort cannot be performed.

## Key Concepts:
-   **Directed Acyclic Graph (DAG):** A directed graph with no directed cycles.
-   **Linear Ordering:** A sequence of all vertices where dependencies are respected.
-   **In-degree:** The number of incoming edges to a vertex.
-   **Out-degree:** The number of outgoing edges from a vertex.

## How Topological Sort Works:

There are two main algorithms for topological sorting:

### 1. DFS-based Algorithm:
-   **Concept:** Perform a Depth-First Search (DFS) traversal on the graph. When a vertex is finished (i.e., all its adjacent vertices have been visited), add it to the front of a list (or push onto a stack).
-   **Steps:**
    1.  Create an empty stack (or list).
    2.  Keep track of visited vertices.
    3.  For each unvisited vertex `u`:
        -   Perform DFS starting from `u`.
        -   During DFS, mark `u` as visited.
        -   Recursively call DFS for all unvisited neighbors of `u`.
        -   After all descendants of `u` have been visited, push `u` onto the stack.
    4.  The topological order is obtained by popping elements from the stack.

### 2. Kahn's Algorithm (BFS-based):
-   **Concept:** Uses the idea of in-degrees. Vertices with an in-degree of 0 have no dependencies and can be placed first in the topological order.
-   **Steps:**
    1.  Compute the in-degree for every vertex.
    2.  Initialize a queue with all vertices that have an in-degree of 0.
    3.  While the queue is not empty:
        -   Dequeue a vertex `u`. Add `u` to the topological order.
        -   For each neighbor `v` of `u`:
            -   Decrement the in-degree of `v`.
            -   If the in-degree of `v` becomes 0, enqueue `v`.
    4.  If the number of vertices in the topological order is less than the total number of vertices, the graph contains a cycle.

## Time Complexity:
-   Both DFS-based and Kahn's algorithms have a time complexity of O(V + E) (where V is the number of vertices and E is the number of edges), as they visit each vertex and edge once.

## Use Cases:
-   **Task Scheduling:** Determining the order in which tasks with dependencies should be executed (e.g., build systems, project management).
-   **Course Scheduling:** Ordering courses based on prerequisites.
-   **Dependency Resolution:** In package managers (e.g., `apt`, `npm`).
-   **Data Serialization:** Ordering objects based on their dependencies.

## Go Example: Topological Sort (DFS-based and Kahn's Algorithm)

See the example code in `graph-topological-sort.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run graph-topological-sort.go
```
