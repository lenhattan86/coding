package main

import (
	"container/list" // For implementing queue for Kahn's algorithm
	"fmt"
)

// --- Graph Representation (Adjacency List for Directed Graph) ---
type DirectedGraph struct {
	numVertices int
	adj         [][]int // Adjacency list
}

func NewDirectedGraph(numV int) *DirectedGraph {
	return &DirectedGraph{
		numVertices: numV,
		adj:         make([][]int, numV),
	}
}

func (g *DirectedGraph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

// --- 1. DFS-based Topological Sort ---

// topologicalSortDFS performs DFS and pushes vertices to stack upon completion
func (g *DirectedGraph) topologicalSortDFS(v int, visited []bool, recursionStack []bool, stack *list.List) bool {
	visited[v] = true
	recursionStack[v] = true

	for _, neighbor := range g.adj[v] {
		if !visited[neighbor] {
			if !g.topologicalSortDFS(neighbor, visited, recursionStack, stack) {
				return false // Cycle detected
			}
		} else if recursionStack[neighbor] {
			return false // Cycle detected
		}
	}
	recursionStack[v] = false // Backtrack
	stack.PushFront(v)        // Push to front of list (simulates stack)
	return true
}

// TopologicalSortDFS performs topological sort using DFS
func (g *DirectedGraph) TopologicalSortDFS() ([]int, bool) {
	visited := make([]bool, g.numVertices)
	recursionStack := make([]bool, g.numVertices) // To detect cycles
	stack := list.New()                           // Use list as a stack

	for i := 0; i < g.numVertices; i++ {
		if !visited[i] {
			if !g.topologicalSortDFS(i, visited, recursionStack, stack) {
				return nil, false // Cycle detected
			}
		}
	}

	result := make([]int, 0, g.numVertices)
	for e := stack.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value.(int))
	}
	return result, true
}

// --- 2. Kahn's Algorithm (BFS-based Topological Sort) ---

// TopologicalSortKahn performs topological sort using Kahn's algorithm
func (g *DirectedGraph) TopologicalSortKahn() ([]int, bool) {
	inDegree := make([]int, g.numVertices)
	for u := 0; u < g.numVertices; u++ {
		for _, v := range g.adj[u] {
			inDegree[v]++
		}
	}

	queue := list.New()
	for i := 0; i < g.numVertices; i++ {
		if inDegree[i] == 0 {
			queue.PushBack(i)
		}
	}

	result := make([]int, 0, g.numVertices)
	countVisited := 0

	for queue.Len() > 0 {
		u := queue.Remove(queue.Front()).(int)
		result = append(result, u)
		countVisited++

		for _, v := range g.adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue.PushBack(v)
			}
		}
	}

	if countVisited != g.numVertices {
		return nil, false // Cycle detected
	}
	return result, true
}

func main() {
	fmt.Println("--- Topological Sort Examples ---")

	// Example 1: A simple DAG
	// 5 -> 0 <- 4
	// |    ^    |
	// v    |    v
	// 2 -> 3 -> 1
	g1 := NewDirectedGraph(6) // Vertices 0 to 5
	g1.AddEdge(5, 2)
	g1.AddEdge(5, 0)
	g1.AddEdge(4, 0)
	g1.AddEdge(4, 1)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 1)

	fmt.Println("\n--- Graph 1 (DAG) ---")
	fmt.Println("DFS-based Topological Sort:")
	if order, ok := g1.TopologicalSortDFS(); ok {
		fmt.Println(order) // Expected: [5 4 2 3 0 1] or similar valid order
	} else {
		fmt.Println("Cycle detected in graph 1!")
	}

	fmt.Println("Kahn's Algorithm (BFS-based) Topological Sort:")
	if order, ok := g1.TopologicalSortKahn(); ok {
		fmt.Println(order) // Expected: [4 5 0 2 3 1] or similar valid order
	} else {
		fmt.Println("Cycle detected in graph 1!")
	}

	// Example 2: Graph with a cycle
	// 0 -> 1 -> 2 -> 0 (cycle)
	g2 := NewDirectedGraph(3)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 0)

	fmt.Println("\n--- Graph 2 (with Cycle) ---")
	fmt.Println("DFS-based Topological Sort:")
	if _, ok := g2.TopologicalSortDFS(); !ok {
		fmt.Println("Cycle detected in graph 2 (DFS)!")
	}

	fmt.Println("Kahn's Algorithm (BFS-based) Topological Sort:")
	if _, ok := g2.TopologicalSortKahn(); !ok {
		fmt.Println("Cycle detected in graph 2 (Kahn's)!")
	}

	fmt.Println("--------------------")
}
