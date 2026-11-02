# String Algorithms: KMP (Knuth-Morris-Pratt) Algorithm

The **Knuth-Morris-Pratt (KMP) algorithm** is an efficient string-searching algorithm that finds occurrences of a "word" (pattern) within a main "text" string. It improves upon the naive string-searching algorithm by avoiding re-checking characters that have already been matched, by using a precomputed table (often called the "LPS array" or "failure function").

## Key Concepts:
-   **Pattern (P):** The string we are searching for.
-   **Text (T):** The main string in which we are searching.
-   **LPS Array (Longest Proper Prefix which is also Suffix):** This is the core of the KMP algorithm. For a given pattern `P`, the `lps[i]` value stores the length of the longest proper prefix of `P[0...i]` that is also a suffix of `P[0...i]`. A "proper prefix" means it's not the entire string itself.

## How KMP Algorithm Works:

### 1. Preprocessing the Pattern (Building the LPS Array):
-   The LPS array helps to know how many characters to shift the pattern when a mismatch occurs, instead of shifting by just one character.
-   It's built by iterating through the pattern and comparing characters.

### 2. Searching the Text:
-   The algorithm compares characters of the pattern with characters of the text.
-   When a match occurs, both text and pattern pointers advance.
-   When a mismatch occurs:
    -   If the pattern pointer is at the beginning (0), simply advance the text pointer.
    -   Otherwise, use the LPS array to shift the pattern. The pattern pointer is moved to `lps[length_of_matched_prefix - 1]`. This means we are effectively aligning the longest proper prefix of the matched part that is also a suffix, and continuing the comparison from there.

## Time Complexity:
-   **Preprocessing (Building LPS Array):** O(m) (where m is the length of the pattern).
-   **Searching:** O(n) (where n is the length of the text).
-   **Overall:** O(n + m). This is a significant improvement over the naive algorithm's O(n*m) worst-case.

## Space Complexity:
-   O(m) for storing the LPS array.

## Use Cases:
-   Efficient string searching in text editors, compilers, and databases.
-   Bioinformatics (e.g., DNA sequence matching).
-   Network intrusion detection systems.

## Go Example: KMP Algorithm Implementation

See the example code in `string-kmp.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run string-kmp.go
```
