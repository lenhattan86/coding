package main

import (
	"fmt"
	"sort"
)

// --- Disjoint Set Union (DSU) Data Structure ---

// DSU represents the Disjoint Set Union data structure
type DSU struct {
	parent []int // parent[i] stores the parent of element i
	rank   []int // rank[i] stores the rank (or height) of the tree rooted at i
}

// NewDSU creates a new DSU structure with n elements
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // Each element is initially its own parent
		rank[i] = 0   // Initial rank is 0
	}
	return &DSU{parent: parent, rank: rank}
}

// Find returns the representative (root) of the set containing element i
// Uses path compression for optimization
func (d *DSU) Find(i int) int {
	if d.parent[i] == i {
		return i
	}
	d.parent[i] = d.Find(d.parent[i]) // Path compression
	return d.parent[i]
}

// Union merges the sets containing elements i and j
// Uses union by rank for optimization
func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)

	if rootI != rootJ {
		// Attach smaller rank tree under root of high rank tree
		if d.rank[rootI] < d.rank[rootJ] {
			d.parent[rootI] = rootJ
		} else if d.rank[rootI] > d.rank[rootJ] {
			d.parent[rootJ] = rootI
		} else {
			d.parent[rootJ] = rootI
			d.rank[rootI]++ // If ranks are same, then increment rank of new root
		}
	}
}

// --- Kruskal's Algorithm ---

// KruskalEdge represents an edge in the graph for Kruskal's algorithm
type KruskalEdge struct {
	u, v   int
	weight int
}

// ByWeight implements sort.Interface for []KruskalEdge based on the weight field.
type ByWeight []KruskalEdge

func (a ByWeight) Len() int           { return len(a) }
func (a ByWeight) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByWeight) Less(i, j int) bool { return a[i].weight < a[j].weight }

// KruskalMST finds the Minimum Spanning Tree of the graph using Kruskal's algorithm
func KruskalMST(numVertices int, edges []KruskalEdge) (int, []KruskalEdge) {
	// 1. Sort all edges in non-decreasing order of their weights
	sort.Sort(ByWeight(edges))

	dsu := NewDSU(numVertices)
	mstEdges := make([]KruskalEdge, 0)
	totalWeight := 0
	edgesCount := 0

	// 2. Iterate through sorted edges
	for _, edge := range edges {
		// If adding this edge does not form a cycle
		if dsu.Find(edge.u) != dsu.Find(edge.v) {
			dsu.Union(edge.u, edge.v) // Merge the two components
			mstEdges = append(mstEdges, edge)
			totalWeight += edge.weight
			edgesCount++

			// Optimization: MST has V-1 edges
			if edgesCount == numVertices-1 {
				break
			}
		}
	}

	// Check if all vertices are connected (for disconnected graphs)
	if edgesCount != numVertices-1 && numVertices > 1 {
		fmt.Println("Graph is disconnected, MST not possible for all vertices.")
		return -1, nil // Indicate disconnected graph
	}

	return totalWeight, mstEdges
}

func main() {
	fmt.Println("--- Kruskal's Algorithm for MST ---")

	// Example Graph (same as Prim's example for comparison):
	//   (0) --4-- (1)
	//   | \     / |
	//   8  7   11 2
	//   |   \ /   |
	//   (7) --1-- (2)
	//   |   / \   |
	//   7  6   14 4
	//   | /     \ |
	//   (6) --9-- (5)
	//    \       /
	//     10 -- 1
	//      \   /
	//       (3)

	numVertices := 9
	edges := []KruskalEdge{
		{0, 1, 4},
		{0, 7, 8},
		{1, 2, 8},
		{1, 7, 11},
		{2, 3, 7},
		{2, 8, 2},
		{2, 5, 4},
		{3, 4, 9},
		{3, 5, 14},
		{4, 5, 10},
		{4, 6, 2},
		{5, 6, 1},
		{6, 7, 1},
		{6, 8, 6},
		{7, 8, 7},
	}

	totalWeight, mstEdges := KruskalMST(numVertices, edges)

	if totalWeight != -1 {
		fmt.Printf("Total weight of MST: %d\n", totalWeight) // Expected: 37
		fmt.Println("Edges in MST:")
		for _, edge := range mstEdges {
			fmt.Printf("  %d - %d (weight %d)\n", edge.u, edge.v, edge.weight)
		}
	}

	fmt.Println("--------------------")
}
