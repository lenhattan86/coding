package main

import "fmt"

// computeLPSArray computes the Longest Proper Prefix which is also Suffix array (LPS array)
// for the given pattern.
func computeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0 // Length of the previous longest prefix suffix
	i := 1

	lps[0] = 0 // lps[0] is always 0

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1] // Don't increment i here
			} else { // if length is 0
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// KMPSearch finds all occurrences of 'pattern' in 'text' using the KMP algorithm.
// Returns a slice of starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return []int{} // Empty pattern found at all positions
	}
	if n == 0 || m > n {
		return []int{} // No text or pattern longer than text
	}

	lps := computeLPSArray(pattern)
	fmt.Printf("Pattern: \"%s\", LPS Array: %v\n", pattern, lps)

	var occurrences []int
	i := 0 // Index for text
	j := 0 // Index for pattern

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			occurrences = append(occurrences, i-j)
			j = lps[j-1] // Shift pattern to find next occurrence
		} else if i < n && pattern[j] != text[i] {
			// Mismatch after j matches
			if j != 0 {
				j = lps[j-1] // Don't match lps[0..lps[j-1]-1] characters,
				// they will match anyway
			} else { // Advance text pointer
				i++
			}
		}
	}
	return occurrences
}

func main() {
	fmt.Println("--- KMP String Searching Algorithm ---")

	text1 := "ABABDABACDABABCABAB"
	pattern1 := "ABABCABAB"
	fmt.Printf("Text: \"%s\"\n", text1)
	fmt.Printf("Pattern: \"%s\"\n", pattern1)
	occurrences1 := KMPSearch(text1, pattern1)
	fmt.Printf("Occurrences found at indices: %v\n", occurrences1) // Expected: [10]
	fmt.Println("--------------------")

	text2 := "AAAAABAAABA"
	pattern2 := "AAAA"
	fmt.Printf("Text: \"%s\"\n", text2)
	fmt.Printf("Pattern: \"%s\"\n", pattern2)
	occurrences2 := KMPSearch(text2, pattern2)
	fmt.Printf("Occurrences found at indices: %v\n", occurrences2) // Expected: [0 1 2 6]
	fmt.Println("--------------------")

	text3 := "ABCABCABC"
	pattern3 := "ABC"
	fmt.Printf("Text: \"%s\"\n", text3)
	fmt.Printf("Pattern: \"%s\"\n", pattern3)
	occurrences3 := KMPSearch(text3, pattern3)
	fmt.Printf("Occurrences found at indices: %v\n", occurrences3) // Expected: [0 3 6]
	fmt.Println("--------------------")

	text4 := "THIS IS A TEST TEXT"
	pattern4 := "TEST"
	fmt.Printf("Text: \"%s\"\n", text4)
	fmt.Printf("Pattern: \"%s\"\n", pattern4)
	occurrences4 := KMPSearch(text4, pattern4)
	fmt.Printf("Occurrences found at indices: %v\n", occurrences4) // Expected: [10]
	fmt.Println("--------------------")

	text5 := "AABAACAADAABAABA"
	pattern5 := "AABA"
	fmt.Printf("Text: \"%s\"\n", text5)
	fmt.Printf("Pattern: \"%s\"\n", pattern5)
	occurrences5 := KMPSearch(text5, pattern5)
	fmt.Printf("Occurrences found at indices: %v\n", occurrences5) // Expected: [0 9 12]
	fmt.Println("--------------------")
}
