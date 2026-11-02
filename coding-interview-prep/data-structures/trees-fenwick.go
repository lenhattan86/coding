package main

import (
	"fmt"
)

// FenwickTree (Binary Indexed Tree)
type FenwickTree struct {
	bit  []int // Binary Indexed Tree array (1-indexed internally)
	size int   // Size of the original array (0-indexed)
}

// NewFenwickTree creates and initializes a new FenwickTree from an array.
func NewFenwickTree(arr []int) *FenwickTree {
	size := len(arr)

ft := &FenwickTree{
		bit:  make([]int, size+1), // BIT is 1-indexed
		size: size,
	}
	for i := 0; i < size; i++ {
		ft.Update(i, arr[i]) // Build by updating each element
	}
	return ft
}

// Update adds 'val' to the element at 'idx' (0-indexed) in the original array.
// This propagates the change through the Fenwick Tree.
func (ft *FenwickTree) Update(idx, val int) {
	idx++ // Convert to 1-indexed for BIT operations
	for idx <= ft.size {
		ft.bit[idx] += val
		idx += idx & (-idx) // Move to the next relevant parent
	}
}

// QueryPrefixSum returns the sum of elements from index 0 to 'idx' (0-indexed) inclusive.
func (ft *FenwickTree) QueryPrefixSum(idx int) int {
	idx++ // Convert to 1-indexed for BIT operations
	sum := 0
	for idx > 0 {
		sum += ft.bit[idx]
		idx -= idx & (-idx) // Move to the next relevant parent
	}
	return sum
}

// QueryRangeSum returns the sum of elements in the range [l, r] (0-indexed).
func (ft *FenwickTree) QueryRangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return ft.QueryPrefixSum(r) - ft.QueryPrefixSum(l-1)
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	fmt.Println("Original Array:", arr)

	ft := NewFenwickTree(arr)
	fmt.Println("Fenwick Tree (internal BIT array):", ft.bit)

	// Query examples
	fmt.Println("\n--- Query Prefix Sum ---")
	fmt.Printf("Prefix sum up to index 0 (arr[0]): %d (Expected: 1)\n", ft.QueryPrefixSum(0))
	fmt.Printf("Prefix sum up to index 2 (arr[0..2]): %d (Expected: 9)\n", ft.QueryPrefixSum(2)) // 1+3+5
	fmt.Printf("Prefix sum up to index 5 (arr[0..5]): %d (Expected: 36)\n", ft.QueryPrefixSum(5)) // 1+3+5+7+9+11

	fmt.Println("\n--- Query Range Sum ---")
	fmt.Printf("Sum of range [1, 3]: %d (Expected: 15)\n", ft.QueryRangeSum(1, 3)) // 3+5+7
	fmt.Printf("Sum of range [2, 2]: %d (Expected: 5)\n", ft.QueryRangeSum(2, 2))   // 5
	fmt.Printf("Sum of range [4, 5]: %d (Expected: 20)\n", ft.QueryRangeSum(4, 5)) // 9+11

	// Update example
	fmt.Println("\n--- Update Element ---")
	updateIndex := 3
	oldValue := arr[updateIndex]
	newValue := 10
	diff := newValue - oldValue
	fmt.Printf("Updating element at index %d from %d to %d (diff: %d)\n", updateIndex, oldValue, newValue, diff)
	ft.Update(updateIndex, diff) // Update with the difference
	arr[updateIndex] = newValue // Also update the original array conceptually
	fmt.Println("Conceptual Updated Array:", arr)
	fmt.Println("Fenwick Tree (internal BIT array after update):", ft.bit)

	// Query after update
	fmt.Println("\n--- Query After Update ---")
	fmt.Printf("Prefix sum up to index 5 (arr[0..5]): %d (Expected: 39)\n", ft.QueryPrefixSum(5)) // 1+3+5+10+9+11
	fmt.Printf("Sum of range [1, 3]: %d (Expected: 18)\n", ft.QueryRangeSum(1, 3)) // 3+5+10
}
