package main

import (
	"container/heap"
	"fmt"
	"math"
)

// --- Dijkstra's Algorithm ---

// Edge represents an edge in the graph
type Edge struct {
	to     int
	weight int
}

// Graph represents the graph using an adjacency list
type Graph struct {
	vertices int
	adj      [][]Edge
}

// NewGraph creates a new graph with V vertices
func NewGraph(V int) *Graph {
	return &Graph{
		vertices: V,
		adj:      make([][]Edge, V),
	}
}

// AddEdge adds a directed edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.adj[from] = append(g.adj[from], Edge{to, weight})
}

// Item is an item in the priority queue
type Item struct {
	value    int // The vertex ID
	priority int // The distance from source
	index    int // The index of the item in the heap
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority (shortest distance)
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest path from a source vertex to all other vertices
func (g *Graph) Dijkstra(source int) []int {
	dist := make([]int, g.vertices)
	for i := range dist {
		dist[i] = math.MaxInt32 // Initialize distances to infinity
	}
	dist[source] = 0 // Distance to source is 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: source, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		u := item.value
		d := item.priority

		// If we found a shorter path to u already, skip
		if d > dist[u] {
			continue
		}

		for _, edge := range g.adj[u] {
			v := edge.to
			weight := edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(&pq, &Item{value: v, priority: dist[v]})
			}
		}
	}
	return dist
}

func main() {
	fmt.Println("--- Dijkstra's Algorithm ---")

	// Example Graph:
	// 0 --10--> 1 --1--> 2
	// |         ^      /|\
	// 5         |       | 
	// |         |       | 
	// v         |       | 
	// 3 --9----> 4 --4-->
	//           |
	//           6
	//           |
	//           v
	//           5 (isolated for this example)

	g := NewGraph(6) // 0-indexed vertices: 0, 1, 2, 3, 4, 5
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 4, 1) // This edge creates a shorter path to 4 from 0 via 1
	g.AddEdge(2, 4, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(4, 5, 6) // Edge to an isolated node 5

	source := 0
	shortestDistances := g.Dijkstra(source)

	fmt.Printf("Shortest distances from source %d:\n", source)
	for i, d := range shortestDistances {
		if d == math.MaxInt32 {
			fmt.Printf("Vertex %d: Unreachable\n", i)
		} else {
			fmt.Printf("Vertex %d: %d\n", i, d)
		}
	}
	// Expected Output:
	// Vertex 0: 0
	// Vertex 1: 10
	// Vertex 2: 11
	// Vertex 3: 5
	// Vertex 4: 11 (via 0->1->4, not 0->3->4 which is 14)
	// Vertex 5: 17

	fmt.Println("--------------------")
}
