package main

import (
	"fmt"
	"math/rand"
	"time"
)

// --- Helper function to swap elements ---
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// --- 1. Bubble Sort ---
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				swapped = true
			}
		}
		if !swapped { // Optimization: if no two elements were swapped by inner loop, array is sorted
			break
		}
	}
}

// --- 2. Selection Sort ---
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// --- 3. Insertion Sort ---
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// --- 4. Merge Sort ---
func MergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mid := n / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// --- 5. Quick Sort ---
func QuickSort(arr []int) {
	quickSortRecursive(arr, 0, len(arr)-1)
}

func quickSortRecursive(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortRecursive(arr, low, pi-1)
		quickSortRecursive(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, high)
	return i + 1
}

// --- 6. Heap Sort ---
// Max-Heap implementation for Heap Sort
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap: i > j
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func HeapSort(arr []int) {
	n := len(arr)
	// Build a max-heap from the input array
	h := (MaxHeap)(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(&h, n, i)
	}

	// One by one extract elements from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		h.Swap(0, i)
		// call max heapify on the reduced heap
		heapify(&h, i, 0)
	}
}

// heapify a subtree rooted with node i which is an index in arr. n is size of heap
func heapify(h *MaxHeap, n, i int) {
	largest := i    // Initialize largest as root
	l := 2*i + 1 // left = 2*i + 1
	r := 2*i + 2 // right = 2*i + 2

	// If left child is larger than root
	if l < n && (*h)[l] > (*h)[largest] {
		largest = l
	}

	// If right child is larger than largest so far
	if r < n && (*h)[r] > (*h)[largest] {
		largest = r
	}

	// If largest is not root
	if largest != i {
		h.Swap(i, largest)
		// Recursively heapify the affected sub-tree
		heapify(h, n, largest)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Helper to create a random array
	createRandomArray := func(size int) []int {
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(100) // Random numbers up to 99
		}
		return arr
	}

	// Test Bubble Sort
	arrBubble := createRandomArray(10)
	fmt.Println("Original (Bubble):", arrBubble)
	BubbleSort(arrBubble)
	fmt.Println("Sorted (Bubble):", arrBubble)
	fmt.Println("--------------------")

	// Test Selection Sort
	arrSelection := createRandomArray(10)
	fmt.Println("Original (Selection):", arrSelection)
	SelectionSort(arrSelection)
	fmt.Println("Sorted (Selection):", arrSelection)
	fmt.Println("--------------------")

	// Test Insertion Sort
	arrInsertion := createRandomArray(10)
	fmt.Println("Original (Insertion):", arrInsertion)
	InsertionSort(arrInsertion)
	fmt.Println("Sorted (Insertion):", arrInsertion)
	fmt.Println("--------------------")

	// Test Merge Sort
	arrMerge := createRandomArray(10)
	fmt.Println("Original (Merge):", arrMerge)
	sortedMerge := MergeSort(arrMerge)
	fmt.Println("Sorted (Merge):", sortedMerge)
	fmt.Println("--------------------")

	// Test Quick Sort
	arrQuick := createRandomArray(10)
	fmt.Println("Original (Quick):", arrQuick)
	QuickSort(arrQuick)
	fmt.Println("Sorted (Quick):", arrQuick)
	fmt.Println("--------------------")

	// Test Heap Sort
	arrHeap := createRandomArray(10)
	fmt.Println("Original (Heap):", arrHeap)
	HeapSort(arrHeap)
	fmt.Println("Sorted (Heap):", arrHeap)
	fmt.Println("--------------------")
}
