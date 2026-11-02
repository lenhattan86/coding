# Trees: Trie (Prefix Tree)

A **Trie**, also known as a **Prefix Tree**, is a tree-like data structure used to store a dynamic set of strings where the keys are usually strings. Unlike a binary search tree, nodes in a trie do not store the key directly. Instead, the position of a node in the tree defines the key with which it is associated. All the descendants of a node have a common prefix of the string associated with that node.

## Characteristics:
-   **Efficient Prefix Matching:** Very efficient for operations like finding all words with a given prefix.
-   **Space Complexity:** Can be high if the dataset contains many unique characters and long strings, as each node might have many children pointers. However, it can be more space-efficient than hash tables for storing many short strings with common prefixes.
-   **Time Complexity:** Operations (insert, search, prefix search) are proportional to the length of the key, not the number of keys.

## Node Structure:
Each node in a trie typically consists of:
-   **`children`:** An array or map of pointers to child nodes, where each pointer corresponds to a character.
-   **`isEndOfWord` (or similar boolean):** A flag indicating whether the path from the root to this node forms a complete word.

## Key Operations:
-   **Insert:** Adds a new word to the trie. It traverses the trie based on the characters of the word, creating new nodes as needed.
-   **Search:** Checks if a given word exists in the trie.
-   **StartsWith (Prefix Search):** Checks if any word in the trie starts with a given prefix.

## Time Complexity:
Let `L` be the length of the key (word).
-   **Insert:** O(L)
-   **Search:** O(L)
-   **StartsWith:** O(L)

## Use Cases:
-   **Autocomplete and Autocorrect:** Suggesting words as a user types.
-   **Spell Checkers:** Efficiently checking if a word is valid.
-   **IP Routing:** Storing IP addresses and finding the longest prefix match.
-   **Dictionary Search:** Fast lookup of words.
-   **T9 predictive text.**

## Go Example: Trie Implementation

See the example code in `trees-trie.go`.
```go
// To run this code:
// 1. Navigate to the data-structures directory: cd coding-interview-prep/data-structures
// 2. Run the file: go run trees-trie.go
```
