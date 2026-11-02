package main

import "fmt"

// --- Basic Bitwise Operations ---

// IsEven checks if a number is even using bitwise AND.
func IsEven(n int) bool {
	return (n & 1) == 0
}

// IsOdd checks if a number is odd using bitwise AND.
func IsOdd(n int) bool {
	return (n & 1) == 1
}

// GetBit returns the k-th bit of n (0-indexed).
func GetBit(n, k int) int {
	return (n >> k) & 1
}

// SetBit sets the k-th bit of n to 1.
func SetBit(n, k int) int {
	return n | (1 << k)
}

// ClearBit clears the k-th bit of n to 0.
func ClearBit(n, k int) int {
	return n & ^(1 << k)
}

// ToggleBit toggles the k-th bit of n.
func ToggleBit(n, k int) int {
	return n ^ (1 << k)
}

// --- Common Bit Manipulation Problems ---

// CountSetBits counts the number of set bits (1s) in a number using Brian Kernighan's algorithm.
func CountSetBits(n int) int {
	count := 0
	for n > 0 {
		n &= (n - 1) // This clears the least significant set bit
		count++
	}
	return count
}

// IsPowerOfTwo checks if a number is a power of two.
func IsPowerOfTwo(n int) bool {
	return (n > 0) && ((n & (n - 1)) == 0)
}

// SwapNumbers swaps two numbers without using a temporary variable.
func SwapNumbers(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func main() {
	fmt.Println("--- Bit Manipulation Examples ---")

	// Basic Operations
	num := 10 // Binary: 1010
	fmt.Printf("Number: %d (Binary: %b)\n", num, num)

	fmt.Printf("Is %d Even? %t\n", num, IsEven(num)) // Expected: true
	fmt.Printf("Is %d Odd? %t\n", num, IsOdd(num))   // Expected: false

	k := 1 // 1st bit (0-indexed)
	fmt.Printf("Bit at position %d: %d\n", k, GetBit(num, k)) // Expected: 1

	newNum := SetBit(num, 0) // Set 0th bit (1010 -> 1011)
	fmt.Printf("Set 0th bit of %d: %d (Binary: %b)\n", num, newNum, newNum) // Expected: 11

	newNum = ClearBit(num, 1) // Clear 1st bit (1010 -> 1000)
	fmt.Printf("Clear 1st bit of %d: %d (Binary: %b)\n", num, newNum, newNum) // Expected: 8

	newNum = ToggleBit(num, 2) // Toggle 2nd bit (1010 -> 1110)
	fmt.Printf("Toggle 2nd bit of %d: %d (Binary: %b)\n", num, newNum, newNum) // Expected: 14
	fmt.Println("--------------------")

	// Common Problems
	fmt.Println("\n--- Common Bit Manipulation Problems ---")
	num = 13 // Binary: 1101
	fmt.Printf("Number: %d (Binary: %b)\n", num, num)
	fmt.Printf("Count set bits in %d: %d\n", num, CountSetBits(num)) // Expected: 3

	fmt.Printf("Is 16 a power of two? %t\n", IsPowerOfTwo(16)) // Expected: true
	fmt.Printf("Is 15 a power of two? %t\n", IsPowerOfTwo(15)) // Expected: false

	a, b := 5, 10
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	a, b = SwapNumbers(a, b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b) // Expected: a = 10, b = 5
	fmt.Println("--------------------")
}
