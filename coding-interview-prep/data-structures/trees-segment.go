package main

import (
	"fmt"
)

// SegmentTree represents a segment tree structure
type SegmentTree struct {
	tree []int // Stores the segment tree nodes
	arr  []int // Original input array
	n    int   // Size of the original array
}

// NewSegmentTree creates and initializes a new SegmentTree
func NewSegmentTree(inputArray []int) *SegmentTree {
	n := len(inputArray)
	// The size of the segment tree array is typically 2*2^ceil(log2(n))
	// A common heuristic is 4*n for safety.
	treeSize := 4 * n
	st := &SegmentTree{
		tree: make([]int, treeSize),
		arr:  inputArray,
		n:    n,
	}
	st.build(0, 0, n-1) // Build the tree from the root (index 0)
	return st
}

// build recursively constructs the segment tree
// nodeIdx: current node's index in the tree array
// start, end: range covered by the current node in the original array
func (st *SegmentTree) build(nodeIdx, start, end int) {
	if start == end { // Leaf node
		st.tree[nodeIdx] = st.arr[start]
	} else {
		mid := (start + end) / 2
		// Recursively build left and right children
		st.build(2*nodeIdx+1, start, mid)
		st.build(2*nodeIdx+2, mid+1, end)
		// Current node's value is the sum of its children
		st.tree[nodeIdx] = st.tree[2*nodeIdx+1] + st.tree[2*nodeIdx+2]
	}
}

// Update updates the value at a specific index in the original array and propagates changes
// nodeIdx: current node's index in the tree array
// start, end: range covered by the current node
// idx: index in the original array to update
// val: new value for arr[idx]
func (st *SegmentTree) Update(idx, val int) {
	st.update(0, 0, st.n-1, idx, val)
}

func (st *SegmentTree) update(nodeIdx, start, end, idx, val int) {
	if start == end { // Leaf node, update the value
		st.arr[idx] = val
		st.tree[nodeIdx] = val
	} else {
		mid := (start + end) / 2
		if start <= idx && idx <= mid { // If idx is in the left child's range
			st.update(2*nodeIdx+1, start, mid, idx, val)
		} else { // If idx is in the right child's range
			st.update(2*nodeIdx+2, mid+1, end, idx, val)
		}
		// Update current node's value based on children
		st.tree[nodeIdx] = st.tree[2*nodeIdx+1] + st.tree[2*nodeIdx+2]
	}
}

// QueryRangeSum queries the sum of elements in a given range [l, r]
// nodeIdx: current node's index in the tree array
// start, end: range covered by the current node
// l, r: query range
func (st *SegmentTree) QueryRangeSum(l, r int) int {
	return st.queryRangeSum(0, 0, st.n-1, l, r)
}

func (st *SegmentTree) queryRangeSum(nodeIdx, start, end, l, r int) int {
	// Case 1: Current node's range is completely outside the query range
	if r < start || end < l {
		return 0
	}

	// Case 2: Current node's range is completely inside the query range
	if l <= start && end <= r {
		return st.tree[nodeIdx]
	}

	// Case 3: Current node's range partially overlaps with the query range
	mid := (start + end) / 2
	p1 := st.queryRangeSum(2*nodeIdx+1, start, mid, l, r)
	p2 := st.queryRangeSum(2*nodeIdx+2, mid+1, end, l, r)
	return p1 + p2
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	fmt.Println("Original Array:", arr)

	st := NewSegmentTree(arr)
	fmt.Println("Segment Tree (internal representation):", st.tree)

	// Query examples
	fmt.Println("\n--- Query Range Sum ---")
	fmt.Printf("Sum of range [0, 5]: %d (Expected: 46)\n", st.QueryRangeSum(0, 5))
	fmt.Printf("Sum of range [1, 3]: %d (Expected: 15)\n", st.QueryRangeSum(1, 3)) // 3+5+7
	fmt.Printf("Sum of range [2, 2]: %d (Expected: 5)\n", st.QueryRangeSum(2, 2))   // 5
	fmt.Printf("Sum of range [4, 5]: %d (Expected: 20)\n", st.QueryRangeSum(4, 5)) // 9+11

	// Update example
	fmt.Println("\n--- Update Element ---")
	updateIndex := 3
	newValue := 10
	fmt.Printf("Updating element at index %d from %d to %d\n", updateIndex, arr[updateIndex], newValue)
	st.Update(updateIndex, newValue)
	fmt.Println("Updated Array:", st.arr)
	fmt.Println("Segment Tree (internal representation after update):", st.tree)

	// Query after update
	fmt.Println("\n--- Query After Update ---")
	fmt.Printf("Sum of range [0, 5]: %d (Expected: 49)\n", st.QueryRangeSum(0, 5)) // 1+3+5+10+9+11
	fmt.Printf("Sum of range [1, 3]: %d (Expected: 18)\n", st.QueryRangeSum(1, 3)) // 3+5+10
}
