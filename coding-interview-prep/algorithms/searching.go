package main

import "fmt"

// --- 1. Linear Search ---
// LinearSearch finds the index of a target element in an array.
// Returns the index if found, -1 otherwise.
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1
}

// --- 2. Binary Search ---
// BinarySearch finds the index of a target element in a SORTED array.
// Returns the index if found, -1 otherwise.
func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2 // To prevent potential overflow for very large low/high
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchRecursive is a recursive implementation of Binary Search.
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchRec(arr, target, 0, len(arr)-1)
}

func binarySearchRec(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRec(arr, target, mid+1, high)
	} else {
		return binarySearchRec(arr, target, low, mid-1)
	}
}

func main() {
	// Test Linear Search
	unsortedArr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Linear Search:")
	fmt.Printf("Array: %v, Target: 22 -> Index: %d\n", unsortedArr, 22, LinearSearch(unsortedArr, 22)) // Expected: 4
	fmt.Printf("Array: %v, Target: 99 -> Index: %d\n", unsortedArr, 99, LinearSearch(unsortedArr, 99)) // Expected: -1
	fmt.Println("--------------------")

	// Test Binary Search (Iterative)
	sortedArr := []int{11, 12, 22, 25, 34, 64, 90}
	fmt.Println("Binary Search (Iterative):")
	fmt.Printf("Array: %v, Target: 25 -> Index: %d\n", sortedArr, 25, BinarySearch(sortedArr, 25)) // Expected: 3
	fmt.Printf("Array: %v, Target: 11 -> Index: %d\n", sortedArr, 11, BinarySearch(sortedArr, 11)) // Expected: 0
	fmt.Printf("Array: %v, Target: 90 -> Index: %d\n", sortedArr, 90, BinarySearch(sortedArr, 90)) // Expected: 6
	fmt.Printf("Array: %v, Target: 33 -> Index: %d\n", sortedArr, 33, BinarySearch(sortedArr, 33)) // Expected: -1
	fmt.Println("--------------------")

	// Test Binary Search (Recursive)
	fmt.Println("Binary Search (Recursive):")
	fmt.Printf("Array: %v, Target: 25 -> Index: %d\n", sortedArr, 25, BinarySearchRecursive(sortedArr, 25)) // Expected: 3
	fmt.Printf("Array: %v, Target: 11 -> Index: %d\n", sortedArr, 11, BinarySearchRecursive(sortedArr, 11)) // Expected: 0
	fmt.Printf("Array: %v, Target: 90 -> Index: %d\n", sortedArr, 90, BinarySearchRecursive(sortedArr, 90)) // Expected: 6
	fmt.Printf("Array: %v, Target: 33 -> Index: %d\n", sortedArr, 33, BinarySearchRecursive(sortedArr, 33)) // Expected: -1
	fmt.Println("--------------------")
}
