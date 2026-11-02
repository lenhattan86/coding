package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32 / 2 // Use a large enough value for infinity to prevent overflow on addition

// FloydWarshall finds the shortest paths between all pairs of vertices
// in a weighted, directed graph.
// It can handle negative edge weights but detects negative cycles.
func FloydWarshall(graph [][]int) [][]int {
	V := len(graph)
	dist := make([][]int, V)
	for i := range dist {
		dist[i] = make([]int, V)
		copy(dist[i], graph[i]) // Initialize dist matrix with input graph
	}

	// Main loops for Floyd-Warshall
	// k is the intermediate vertex
	for k := 0; k < V; k++ {
		// i is the source vertex
		for i := 0; i < V; i++ {
			// j is the destination vertex
			for j := 0; j < V; j++ {
				// If vertex k is on the shortest path from i to j,
				// then update the value of dist[i][j]
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Check for negative cycles
	for i := 0; i < V; i++ {
		if dist[i][i] < 0 {
			fmt.Println("Graph contains a negative cycle!")
			// Optionally, you can return a special value or error
			// For this example, we'll just print and continue with potentially incorrect paths
		}
	}

	return dist
}

func main() {
	fmt.Println("--- Floyd-Warshall Algorithm ---")

	// Example Graph (Adjacency Matrix representation)
	// INF represents no direct edge
	// V = 4
	// Graph:
	//      (0) --3--> (1)
	//       ^ \       /
	//       |  \     /
	//       6   5   -2
	//       |    \ /
	//      (3) <--4-- (2)
	//
	// Initial distance matrix:
	//   0  1  2  3
	// 0 [0, 3, INF, INF]
	// 1 [INF, 0, -2, INF]
	// 2 [INF, INF, 0, 4]
	// 3 [6, INF, INF, 0]

	graph := [][]int{
		{0, 3, INF, INF},
		{INF, 0, -2, INF},
		{INF, INF, 0, 4},
		{6, INF, INF, 0},
	}

	shortestPaths := FloydWarshall(graph)

	fmt.Println("Shortest path distances between all pairs of vertices:")
	for i := range shortestPaths {
		for j := range shortestPaths[i] {
			if shortestPaths[i][j] == INF {
				fmt.Printf("%7s", "INF")
			} else {
				fmt.Printf("%7d", shortestPaths[i][j])
			}
		}
		fmt.Println()
	}
	// Expected Output:
	//       0      1      2      3
	// 0     0      3      1      5
	// 1   INF      0     -2      2
	// 2   INF    INF      0      4
	// 3     6      9      7     11 (if path 3->0->1->2->3 is considered, but 3->0->1->2 is 6+3-2=7, 3->0->1->2->3 is 7+4=11)

	fmt.Println("--------------------\n")

	// Example with a negative cycle (0 -> 1 -> 2 -> 0 with total weight -1)
	fmt.Println("--- Example with Negative Cycle ---")
	graphWithNegativeCycle := [][]int{
		{0, 1, INF},
		{INF, 0, -2},
		{-1, INF, 0},
	}
	shortestPathsNegativeCycle := FloydWarshall(graphWithNegativeCycle)
	fmt.Println("Shortest path distances (with potential negative cycle):")
	for i := range shortestPathsNegativeCycle {
		for j := range shortestPathsNegativeCycle[i] {
			if shortestPathsNegativeCycle[i][j] == INF {
				fmt.Printf("%7s", "INF")
			} else {
				fmt.Printf("%7d", shortestPathsNegativeCycle[i][j])
			}
		}
		fmt.Println()
	}
	// Expected to print "Graph contains a negative cycle!" and then show results
	// where dist[i][i] might be negative.
}
