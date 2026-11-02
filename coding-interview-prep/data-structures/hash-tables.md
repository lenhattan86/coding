# Hash Tables (Hash Maps)

A hash table (also known as a hash map) is a data structure that implements an associative array abstract data type, a structure that can map keys to values. A hash table uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

## Key Concepts:
-   **Hash Function:** A function that converts a key into an integer index (hash value) within the array of buckets. A good hash function distributes keys uniformly across the buckets to minimize collisions.
-   **Collision:** Occurs when the hash function generates the same index for two or more different keys.
-   **Collision Resolution:** Techniques to handle collisions. Common methods include:
    -   **Separate Chaining:** Each bucket is a linked list (or another data structure) that stores all elements that hash to that bucket.
    -   **Open Addressing:** When a collision occurs, the system probes for another empty slot in the array.
        -   **Linear Probing:** Search for the next empty slot sequentially.
        -   **Quadratic Probing:** Search for the next empty slot using a quadratic function.
        -   **Double Hashing:** Use a second hash function to determine the step size for probing.
-   **Load Factor:** The ratio of the number of items stored in the hash table to the number of buckets. A high load factor increases the probability of collisions and degrades performance. When the load factor exceeds a certain threshold, the hash table is typically resized (rehashed).

## Characteristics:
-   Provides average O(1) time complexity for insertion, deletion, and search operations.
-   Worst-case time complexity can be O(n) if all keys hash to the same bucket (due to poor hash function or many collisions).
-   Requires careful design of hash function and collision resolution strategy.

## Time Complexity:
-   **Average Case:**
    -   **Insert:** O(1)
    -   **Search:** O(1)
    -   **Delete:** O(1)
-   **Worst Case:**
    -   **Insert:** O(n)
    -   **Search:** O(n)
    -   **Delete:** O(n)

## Use Cases:
-   Symbol tables in compilers.
-   Database indexing.
-   Caches.
-   Associative arrays (maps in Go).
-   Set implementations.

## Go Example: Hash Table Implementation (with Separate Chaining)

See the example code in `hash-tables.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run hash-tables.go
```
