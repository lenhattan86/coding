package main

import "fmt"

const ALPHABET_SIZE = 26 // For lowercase English letters 'a' through 'z'

// TrieNode represents a node in the Trie
type TrieNode struct {
	children [ALPHABET_SIZE]*TrieNode
	isEndOfWord bool
}

// NewTrieNode creates and returns a new TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

// Trie represents the Trie data structure
type Trie struct {
	root *TrieNode
}

// NewTrie creates and returns a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		index := char - 'a' // Assuming lowercase English letters
		if current.children[index] == nil {
			current.children[index] = NewTrieNode()
		}
		current = current.children[index]
	}
	current.isEndOfWord = true
}

// Search checks if a word exists in the Trie
func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		index := char - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current != nil && current.isEndOfWord
}

// StartsWith checks if there is any word in the Trie that starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		index := char - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	return current != nil
}

func main() {
	trie := NewTrie()

	wordsToInsert := []string{"apple", "app", "apricot", "banana", "band", "bat"}
	fmt.Println("Inserting words:", wordsToInsert)
	for _, word := range wordsToInsert {
		trie.Insert(word)
	}

	fmt.Println("\n--- Search Operations ---")
	fmt.Printf("Search 'apple': %t\n", trie.Search("apple"))   // Expected: true
	fmt.Printf("Search 'app': %t\n", trie.Search("app"))     // Expected: true
	fmt.Printf("Search 'ap': %t\n", trie.Search("ap"))      // Expected: false (not a full word)
	fmt.Printf("Search 'apricot': %t\n", trie.Search("apricot")) // Expected: true
	fmt.Printf("Search 'banana': %t\n", trie.Search("banana")) // Expected: true
	fmt.Printf("Search 'band': %t\n", trie.Search("band"))   // Expected: true
	fmt.Printf("Search 'bat': %t\n", trie.Search("bat"))     // Expected: true
	fmt.Printf("Search 'bana': %t\n", trie.Search("bana"))   // Expected: false
	fmt.Printf("Search 'orange': %t\n", trie.Search("orange")) // Expected: false

	fmt.Println("\n--- StartsWith Operations ---")
	fmt.Printf("StartsWith 'ap': %t\n", trie.StartsWith("ap"))     // Expected: true
	fmt.Printf("StartsWith 'app': %t\n", trie.StartsWith("app"))    // Expected: true
	fmt.Printf("StartsWith 'ban': %t\n", trie.StartsWith("ban"))    // Expected: true
	fmt.Printf("StartsWith 'bat': %t\n", trie.StartsWith("bat"))    // Expected: true
	fmt.Printf("StartsWith 'or': %t\n", trie.StartsWith("or"))     // Expected: false
	fmt.Printf("StartsWith 'applepie': %t\n", trie.StartsWith("applepie")) // Expected: false (no word starts with this long prefix)
}
