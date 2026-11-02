package main

import (
	"container/heap"
	"fmt"
	"math"
)

// --- Prim's Algorithm ---

// Edge represents an edge in the graph with a weight
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

// AddEdge adds an undirected edge to the graph
func (g *Graph) AddEdge(u, v, weight int) {
	g.adj[u] = append(g.adj[u], Edge{to: v, weight: weight})
	g.adj[v] = append(g.adj[v], Edge{to: u, weight: weight}) // Undirected
}

// PriorityQueueItem represents an item in the priority queue for Prim's
type PriorityQueueItem struct {
	vertex   int // The vertex
	priority int // The key (minimum weight to connect to MST)
	index    int // The index of the item in the heap
}

// PriorityQueue implements heap.Interface for Prim's algorithm
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
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

// PrimMST finds the Minimum Spanning Tree of the graph using Prim's algorithm
func (g *Graph) PrimMST() (int, []int) {
	key := make([]int, g.vertices)
	parent := make([]int, g.vertices)
	inMST := make([]bool, g.vertices)

	for i := range key {
		key[i] = math.MaxInt32
		parent[i] = -1
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	key[0] = 0
	heap.Push(&pq, &PriorityQueueItem{vertex: 0, priority: 0})

	totalWeight := 0

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityQueueItem)
		u := item.vertex

		if inMST[u] {
			continue
		}

		inMST[u] = true
		totalWeight += key[u]

		for _, edge := range g.adj[u] {
			v := edge.to
			weight := edge.weight

			if !inMST[v] && weight < key[v] {
				key[v] = weight
				parent[v] = u
				heap.Push(&pq, &PriorityQueueItem{vertex: v, priority: key[v]})
			}
		}
	}

	for i := 0; i < g.vertices; i++ {
		if !inMST[i] {
			fmt.Println("Graph is disconnected, MST not possible for all vertices.")
			return -1, nil
		}
	}

	return totalWeight, parent
}

func main() {
	fmt.Println("--- Prim's Algorithm for MST ---")

	g := NewGraph(9)
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 7, 8)
	g.AddEdge(1, 2, 8)
	g.AddEdge(1, 7, 11)
	g.AddEdge(2, 3, 7)
	g.AddEdge(2, 8, 2)
	g.AddEdge(2, 5, 4)
	g.AddEdge(3, 4, 9)
	g.AddEdge(3, 5, 14)
	g.AddEdge(4, 5, 10)
	g.AddEdge(4, 6, 2)
	g.AddEdge(5, 6, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(6, 8, 6)
	g.AddEdge(7, 8, 7)

	totalWeight, parent := g.PrimMST()

	if totalWeight != -1 {
		fmt.Printf("Total weight of MST: %d\n", totalWeight)
		fmt.Println("Edges in MST:")
		for i := 1; i < g.vertices; i++ {
			fmt.Printf("  %d - %d (weight %d)\n", parent[i], i, key[i])
		}
	}

	fmt.Println("--------------------")
}
