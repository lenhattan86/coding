# Algorithms: Bit Manipulation

Bit manipulation is the act of algorithmically manipulating bits or other small pieces of data. It's a powerful technique that can often lead to highly optimized and efficient solutions, especially in competitive programming and low-level systems. Understanding bitwise operations is crucial for many interview problems.

## Key Concepts:
-   **Bits:** The smallest unit of data in a computer, represented as 0 or 1.
-   **Binary Representation:** Numbers are stored in computers as sequences of bits.
-   **Bitwise Operators:** Operations that work directly on the individual bits of a number.

## Bitwise Operators in Go:

1.  **AND (`&`):** Sets each bit to 1 if both corresponding bits are 1.
    -   `1 & 1 = 1`
    -   `1 & 0 = 0`
    -   `0 & 1 = 0`
    -   `0 & 0 = 0`
    -   **Use Case:** Checking if a bit is set, clearing specific bits.

2.  **OR (`|`):** Sets each bit to 1 if at least one of the corresponding bits is 1.
    -   `1 | 1 = 1`
    -   `1 | 0 = 1`
    -   `0 | 1 = 1`
    -   `0 | 0 = 0`
    -   **Use Case:** Setting specific bits.

3.  **XOR (`^`):** Sets each bit to 1 if only one of the corresponding bits is 1 (i.e., if bits are different).
    -   `1 ^ 1 = 0`
    -   `1 ^ 0 = 1`
    -   `0 ^ 1 = 1`
    -   `0 ^ 0 = 0`
    -   **Use Case:** Toggling bits, finding the unique element in an array where others appear twice, swapping numbers without a temporary variable.

4.  **NOT (`^` or `~` in some languages, but `^` for bitwise NOT in Go with a mask):** Inverts all the bits. In Go, `^x` is bitwise complement (XOR with all 1s mask).
    -   `^0 = 1`
    -   `^1 = 0`
    -   **Use Case:** Inverting flags.

5.  **Left Shift (`<<`):** Shifts bits to the left, filling with zeros on the right. Equivalent to multiplying by powers of 2.
    -   `x << n` is `x * 2^n`
    -   **Use Case:** Creating masks, multiplying by powers of 2.

6.  **Right Shift (`>>`):** Shifts bits to the right. For unsigned integers, fills with zeros on the left. For signed integers, fills with the sign bit (arithmetic right shift) or zeros (logical right shift). Go's `>>` is arithmetic for signed and logical for unsigned.
    -   `x >> n` is `x / 2^n` (integer division)
    -   **Use Case:** Dividing by powers of 2, extracting specific bits.

## Common Bit Manipulation Techniques & Problems:

-   **Check if a number is even or odd:** `(n & 1) == 0` for even, `(n & 1) == 1` for odd.
-   **Check if a bit is set:** `(n >> k) & 1 == 1` or `(n & (1 << k)) != 0`.
-   **Set a bit:** `n | (1 << k)`.
-   **Clear a bit:** `n & ^(1 << k)`.
-   **Toggle a bit:** `n ^ (1 << k)`.
-   **Count set bits (Hamming weight):** Various algorithms, including Brian Kernighan's algorithm (`n = n & (n-1)`).
-   **Power of 2 check:** `(n > 0) && ((n & (n - 1)) == 0)`.
-   **Swapping two numbers:** `a = a ^ b; b = a ^ b; a = a ^ b;`

## Go Example: Common Bit Manipulation Operations

See the example code in `bit-manipulation.go`.
```go
// To run this code:
// 1. Navigate to the algorithms directory: cd coding-interview-prep/algorithms
// 2. Run the file: go run bit-manipulation.go
```
