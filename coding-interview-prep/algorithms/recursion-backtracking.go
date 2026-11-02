package main

import "fmt"

// --- Recursion Example: Factorial ---
// Factorial calculates the factorial of a non-negative integer n.
// n! = n * (n-1) * ... * 1
// Base case: 0! = 1
func Factorial(n int) int {
	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return -1 // Or panic, or return an error
	}
	if n == 0 { // Base case
		return 1
	}
	return n * Factorial(n-1) // Recursive step
}

// --- Backtracking Example: N-Queens Problem ---
// The N-Queens puzzle is the problem of placing N chess queens on an N×N chessboard
// so that no two queens threaten each other.
// This means no two queens share the same row, column, or diagonal.

var solutions [][]string // Stores all valid board configurations

// SolveNQueens initiates the N-Queens problem solving
func SolveNQueens(n int) [][]string {
	solutions = [][]string{} // Reset solutions for each call
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.' // Initialize empty board
		}
	}
	placeQueen(board, 0, n)
	return solutions
}

// placeQueen is the recursive backtracking function
func placeQueen(board [][]rune, row, n int) {
	if row == n { // Base case: all queens are placed
		addSolution(board)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(board, row, col, n) {
			board[row][col] = 'Q' // Place queen
			placeQueen(board, row+1, n) // Recurse to next row
			board[row][col] = '.' // Backtrack: remove queen
		}
	}
}

// isValid checks if placing a queen at (row, col) is valid
func isValid(board [][]rune, row, col, n int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// addSolution converts the rune board to a string slice and adds to solutions
func addSolution(board [][]rune) {
	solution := make([]string, len(board))
	for i, row := range board {
		solution[i] = string(row)
	}
	solutions = append(solutions, solution)
}

func main() {
	// Test Factorial
	fmt.Println("--- Factorial ---")
	fmt.Printf("Factorial of 5: %d\n", Factorial(5)) // Expected: 120
	fmt.Printf("Factorial of 0: %d\n", Factorial(0)) // Expected: 1
	fmt.Printf("Factorial of -3: %d\n", Factorial(-3)) // Expected: -1 (with error message)
	fmt.Println("--------------------")

	// Test N-Queens
	fmt.Println("--- N-Queens Problem ---")
	n := 4 // For a 4x4 board
	fmt.Printf("Solving N-Queens for N = %d\n", n)
	nQueensSolutions := SolveNQueens(n)
	for i, sol := range nQueensSolutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range sol {
			fmt.Println(row)
		}
		fmt.Println()
	}
	fmt.Printf("Total solutions for N = %d: %d\n", n, len(nQueensSolutions)) // Expected for N=4: 2 solutions
	fmt.Println("--------------------")

	n = 8 // For an 8x8 board
	fmt.Printf("Solving N-Queens for N = %d\n", n)
	nQueensSolutions = SolveNQueens(n)
	fmt.Printf("Total solutions for N = %d: %d\n", n, len(nQueensSolutions)) // Expected for N=8: 92 solutions
	fmt.Println("--------------------")
}
