package main

import (
	"fmt"
)

const (
	// Initial size of the hash table array
	initialBucketSize = 10
)

// Entry represents a key-value pair in the hash table
type Entry struct {
	key   string
	value interface{}
	next  *Entry // For separate chaining
}

// HashTable represents the hash table data structure
type HashTable struct {
	buckets []*Entry
	size    int // Number of key-value pairs currently in the table
}

// NewHashTable creates and returns a new HashTable
func NewHashTable() *HashTable {
	return &HashTable{
		buckets: make([]*Entry, initialBucketSize),
		size:    0,
	}
}

// hashFunction computes the hash value for a given key
func (ht *HashTable) hashFunction(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % len(ht.buckets) // Simple polynomial rolling hash
	}
	return hash
}

// Insert adds a key-value pair to the hash table
func (ht *HashTable) Insert(key string, value interface{}) {
	index := ht.hashFunction(key)
	newEntry := &Entry{key: key, value: value, next: nil}

	if ht.buckets[index] == nil {
		ht.buckets[index] = newEntry
	} else {
		// Handle collision using separate chaining
		current := ht.buckets[index]
		for current.next != nil {
			// If key already exists, update its value
			if current.key == key {
				current.value = value
				return
			}
			current = current.next
		}
		// Check the last node in the chain
		if current.key == key {
			current.value = value
			return
		}
		current.next = newEntry
	}
	ht.size++
}

// Search retrieves the value associated with a given key
func (ht *HashTable) Search(key string) (interface{}, bool) {
	index := ht.hashFunction(key)
	current := ht.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return nil, false // Key not found
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key string) bool {
	index := ht.hashFunction(key)
	current := ht.buckets[index]
	var prev *Entry

	for current != nil {
		if current.key == key {
			if prev == nil {
				// Deleting the head of the linked list
				ht.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			ht.size--
			return true
		}
		prev = current
		current = current.next
	}
	return false // Key not found
}

// Size returns the number of elements in the hash table
func (ht *HashTable) Size() int {
	return ht.size
}

func main() {
	ht := NewHashTable()
	fmt.Println("Initial hash table size:", ht.Size()) // Expected: 0

	ht.Insert("apple", 10)
	ht.Insert("banana", 20)
	ht.Insert("cherry", 30)
	fmt.Println("Hash table size after inserts:", ht.Size()) // Expected: 3

	value, found := ht.Search("banana")
	if found {
		fmt.Println("Found banana:", value) // Expected: 20
	} else {
		fmt.Println("Banana not found.")
	}

	value, found = ht.Search("grape")
	if found {
		fmt.Println("Found grape:", value)
	} else {
		fmt.Println("Grape not found.") // Expected: Grape not found.
	}

	ht.Insert("apple", 15) // Update existing key
	value, found = ht.Search("apple")
	if found {
		fmt.Println("Updated apple:", value) // Expected: 15
	}

	fmt.Println("Deleting banana...")
	deleted := ht.Delete("banana")
	fmt.Println("Banana deleted:", deleted) // Expected: true
	fmt.Println("Hash table size after delete:", ht.Size()) // Expected: 2

	_, found = ht.Search("banana")
	if found {
		fmt.Println("Banana still found (error).")
	} else {
		fmt.Println("Banana not found after deletion.") // Expected: Banana not found after deletion.
	}

	fmt.Println("Deleting non-existent key 'grape'...")
	deleted = ht.Delete("grape")
	fmt.Println("Grape deleted:", deleted) // Expected: false
}
