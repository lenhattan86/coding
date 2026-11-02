package main

import (
	"fmt"
	"math"
)

// --- Fibonacci Sequence ---

// FibonacciMemoization calculates the nth Fibonacci number using memoization (top-down DP).
var fibMemo map[int]int

func FibonacciMemoization(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}
	if val, ok := fibMemo[n]; ok {
		return val
	}

	result := FibonacciMemoization(n-1) + FibonacciMemoization(n-2)
	fibMemo[n] = result
	return result
}

// FibonacciTabulation calculates the nth Fibonacci number using tabulation (bottom-up DP).
func FibonacciTabulation(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// --- Longest Common Subsequence (LCS) ---
// LCS finds the length of the longest common subsequence of two strings.
// A subsequence is a sequence that can be derived from another sequence by deleting some or no elements
// without changing the order of the remaining elements.

func LongestCommonSubsequence(text1, text2 string) int {
	m := len(text1)
	n := len(text2)

	// dp[i][j] stores the length of LCS of text1[0...i-1] and text2[0...j-1]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}
	return dp[m][n]
}

func main() {
	// Test Fibonacci Memoization
	fmt.Println("--- Fibonacci (Memoization) ---")
	fibMemo = make(map[int]int) // Initialize memoization map
	fmt.Printf("Fibonacci(10): %d\n", FibonacciMemoization(10)) // Expected: 55
	fmt.Printf("Fibonacci(0): %d\n", FibonacciMemoization(0))   // Expected: 0
	fmt.Printf("Fibonacci(1): %d\n", FibonacciMemoization(1))   // Expected: 1
	fmt.Printf("Fibonacci(20): %d\n", FibonacciMemoization(20)) // Expected: 6765
	fmt.Println("--------------------")

	// Test Fibonacci Tabulation
	fmt.Println("--- Fibonacci (Tabulation) ---")
	fmt.Printf("Fibonacci(10): %d\n", FibonacciTabulation(10)) // Expected: 55
	fmt.Printf("Fibonacci(0): %d\n", FibonacciTabulation(0))   // Expected: 0
	fmt.Printf("Fibonacci(1): %d\n", FibonacciTabulation(1))   // Expected: 1
	fmt.Printf("Fibonacci(20): %d\n", FibonacciTabulation(20)) // Expected: 6765
	fmt.Println("--------------------")

	// Test Longest Common Subsequence
	fmt.Println("--- Longest Common Subsequence ---")
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	fmt.Printf("LCS of \"%s\" and \"%s\": %d\n", s1, s2, LongestCommonSubsequence(s1, s2)) // Expected: 4 (GTAB)

	s3 := "ABCBDAB"
	s4 := "BDCABA"
	fmt.Printf("LCS of \"%s\" and \"%s\": %d\n", s3, s4, LongestCommonSubsequence(s3, s4)) // Expected: 4 (BCBA or BDAB)

	s5 := "abc"
	s6 := "abc"
	fmt.Printf("LCS of \"%s\" and \"%s\": %d\n", s5, s6, LongestCommonSubsequence(s5, s6)) // Expected: 3

	s7 := "abc"
	s8 := "def"
	fmt.Printf("LCS of \"%s\" and \"%s\": %d\n", s7, s8, LongestCommonSubsequence(s7, s8)) // Expected: 0
	fmt.Println("--------------------")
}
