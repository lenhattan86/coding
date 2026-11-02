package main

import (
	"container/list" // For implementing queue for BFS
	"fmt"
)

// Graph represents a graph using an adjacency list
type Graph struct {
	numVertices int
	adjList     map[int][]int
}

// NewGraph creates a new graph with the given number of vertices
func NewGraph(numV int) *Graph {
	return &Graph{
		numVertices: numV,
		adjList:     make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph (undirected)
func (g *Graph) AddEdge(v1, v2 int) {
	g.adjList[v1] = append(g.adjList[v1], v2)
	g.adjList[v2] = append(g.adjList[v2], v1) // For undirected graph
}

// BFS performs Breadth-First Search starting from a given vertex
func (g *Graph) BFS(startVertex int) {
	fmt.Printf("BFS starting from vertex %d: ", startVertex)
	visited := make(map[int]bool)
	queue := list.New()

	visited[startVertex] = true
	queue.PushBack(startVertex)

	for queue.Len() > 0 {
		v := queue.Remove(queue.Front()).(int)
		fmt.Printf("%d ", v)

		for _, neighbor := range g.adjList[v] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
	fmt.Println()
}

// DFS performs Depth-First Search starting from a given vertex
func (g *Graph) DFS(startVertex int) {
	fmt.Printf("DFS starting from vertex %d: ", startVertex)
	visited := make(map[int]bool)
	g.dfsRecursive(startVertex, visited)
	fmt.Println()
}

func (g *Graph) dfsRecursive(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Printf("%d ", v)

	for _, neighbor := range g.adjList[v] {
		if !visited[neighbor] {
			g.dfsRecursive(neighbor, visited)
		}
	}
}

// Adjacency Matrix representation (for demonstration)
type AdjacencyMatrixGraph struct {
	numVertices int
	matrix      [][]int
}

// NewAdjacencyMatrixGraph creates a new graph using an adjacency matrix
func NewAdjacencyMatrixGraph(numV int) *AdjacencyMatrixGraph {
	matrix := make([][]int, numV)
	for i := range matrix {
		matrix[i] = make([]int, numV)
	}
	return &AdjacencyMatrixGraph{
		numVertices: numV,
		matrix:      matrix,
	}
}

// AddEdge adds an edge to the adjacency matrix graph (undirected)
func (g *AdjacencyMatrixGraph) AddEdge(v1, v2 int) {
	if v1 >= 0 && v1 < g.numVertices && v2 >= 0 && v2 < g.numVertices {
		g.matrix[v1][v2] = 1
		g.matrix[v2][v1] = 1 // For undirected graph
	}
}

// PrintMatrix prints the adjacency matrix
func (g *AdjacencyMatrixGraph) PrintMatrix() {
	fmt.Println("Adjacency Matrix:")
	for i := 0; i < g.numVertices; i++ {
		for j := 0; j < g.numVertices; j++ {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Example using Adjacency List
	fmt.Println("--- Adjacency List Graph ---")
	graph := NewGraph(5) // 0-indexed vertices: 0, 1, 2, 3, 4
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	fmt.Println("Graph Adjacency List:")
	for v, neighbors := range graph.adjList {
		fmt.Printf("%d: %v\n", v, neighbors)
	}

	graph.BFS(0) // Expected: 0 1 2 3 4
	graph.DFS(0) // Expected: 0 1 3 4 2 (order might vary based on map iteration)

	// Example using Adjacency Matrix
	fmt.Println("\n--- Adjacency Matrix Graph ---")
	matrixGraph := NewAdjacencyMatrixGraph(5)
	matrixGraph.AddEdge(0, 1)
	matrixGraph.AddEdge(0, 2)
	matrixGraph.AddEdge(1, 3)
	matrixGraph.AddEdge(2, 4)
	matrixGraph.AddEdge(3, 4)
	matrixGraph.PrintMatrix()
}
