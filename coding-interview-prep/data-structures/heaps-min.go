package main

import (
	"container/heap" // Go's built-in heap interface
	"fmt"
)

// IntHeap is a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap: i < j
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the minimum element from the heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek returns the minimum element without removing it
func (h IntHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h[0], true
}

func main() {
	// Create a new min-heap
	minHeap := &IntHeap{}
	heap.Init(minHeap) // Initialize the heap

	fmt.Println("Is heap empty?", minHeap.Len() == 0) // Expected: true

	// Insert elements
	fmt.Println("Inserting elements: 5, 3, 8, 1, 7, 2")
	heap.Push(minHeap, 5)
	heap.Push(minHeap, 3)
	heap.Push(minHeap, 8)
	heap.Push(minHeap, 1)
	heap.Push(minHeap, 7)
	heap.Push(minHeap, 2)

	fmt.Printf("Heap elements (internal array, not sorted): %v\n", *minHeap)
	fmt.Println("Heap size:", minHeap.Len()) // Expected: 6

	// Peek at the minimum element
	minVal, ok := minHeap.Peek()
	if ok {
		fmt.Println("Peek min element:", minVal) // Expected: 1
	}

	// Extract elements
	fmt.Println("\nExtracting elements (should be in ascending order):")
	for minHeap.Len() > 0 {
		fmt.Printf("Extracted: %d\n", heap.Pop(minHeap))
	}

	fmt.Println("Is heap empty?", minHeap.Len() == 0) // Expected: true

	// Example with more complex types (e.g., structs) would require custom Less function
	// For simplicity, this example uses int.
}
