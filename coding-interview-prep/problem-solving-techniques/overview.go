package main

import (
	"fmt"
	"sort"
)

// --- Conceptual Example: Two Sum Problem ---
// Problem: Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// Technique 1: Brute Force (O(n^2))
func twoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// Technique 2: Using Hashing (O(n) average)
func twoSumHashing(nums []int, target int) []int {
	numMap := make(map[int]int) // value -> index
	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

// Technique 3: Two Pointers (O(n log n) due to sorting, then O(n))
// This requires the array to be sorted.
func twoSumTwoPointers(nums []int, target int) []int {
	// Create a copy and sort it to preserve original indices if needed
	type pair struct {
		val int
		idx int
	}
	pairs := make([]pair, len(nums))
	for i, num := range nums {
		pairs[i] = pair{val: num, idx: i}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	left, right := 0, len(pairs)-1
	for left < right {
		sum := pairs[left].val + pairs[right].val
		if sum == target {
			return []int{pairs[left].idx, pairs[right].idx}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	fmt.Println("--- Conceptual Problem Solving Techniques ---")

	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("Input: %v, Target: %d\n", nums, target)

	// Brute Force
	resultBF := twoSumBruteForce(nums, target)
	fmt.Printf("Brute Force Result: %v\n", resultBF) // Expected: [0 1]

	// Hashing
	resultHash := twoSumHashing(nums, target)
	fmt.Printf("Hashing Result: %v\n", resultHash) // Expected: [0 1]

	// Two Pointers
	resultTP := twoSumTwoPointers(nums, target)
	fmt.Printf("Two Pointers Result: %v\n", resultTP) // Expected: [0 1] (original indices)

	fmt.Println("\n--- This example demonstrates different algorithmic techniques for the same problem. ---")
	fmt.Println("--- The choice of technique impacts time and space complexity. ---")

	fmt.Println("\n--- Problem Solving Techniques Example Finished ---")
}
