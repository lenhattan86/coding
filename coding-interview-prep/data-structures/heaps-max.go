package main

import (
	"container/heap" // Go's built-in heap interface
	"fmt"
)

// IntMaxHeap is a max-heap of integers.
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap: i > j
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap
func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop removes and returns the maximum element from the heap
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Peek returns the maximum element without removing it
func (h IntMaxHeap) Peek() (int, bool) {
	if h.Len() == 0 {
		return 0, false
	}
	return h[0], true
}

func main() {
	// Create a new max-heap
	maxHeap := &IntMaxHeap{}
	heap.Init(maxHeap) // Initialize the heap

	fmt.Println("Is max-heap empty?", maxHeap.Len() == 0) // Expected: true

	// Insert elements
	fmt.Println("Inserting elements: 5, 3, 8, 1, 7, 2")
	heap.Push(maxHeap, 5)
	heap.Push(maxHeap, 3)
	heap.Push(maxHeap, 8)
	heap.Push(maxHeap, 1)
	heap.Push(maxHeap, 7)
	heap.Push(maxHeap, 2)

	fmt.Printf("Max-Heap elements (internal array, not sorted): %v\n", *maxHeap)
	fmt.Println("Max-Heap size:", maxHeap.Len()) // Expected: 6

	// Peek at the maximum element
	maxVal, ok := maxHeap.Peek()
	if ok {
		fmt.Println("Peek max element:", maxVal) // Expected: 8
	}

	// Extract elements
	fmt.Println("\nExtracting elements (should be in descending order):")
	for maxHeap.Len() > 0 {
		fmt.Printf("Extracted: %d\n", heap.Pop(maxHeap))
	}

	fmt.Println("Is max-heap empty?", maxHeap.Len() == 0) // Expected: true
}
