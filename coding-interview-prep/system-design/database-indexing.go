package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Record represents a simple data record.
type Record struct {
	ID   int
	Name string
	Age  int
}

// generateRecords creates a slice of random records.
func generateRecords(count int) []Record {
	records := make([]Record, count)
	for i := 0; i < count; i++ {
		records[i] = Record{
			ID:   i,
			Name: fmt.Sprintf("User%%d", i),
			Age:  rand.Intn(80) + 18, // Age between 18 and 97
		}
	}
	return records
}

// --- Unindexed Search ---
// LinearSearchRecords performs a linear search on a slice of records.
func LinearSearchRecords(records []Record, targetID int) (*Record, bool) {
	for _, record := range records {
		if record.ID == targetID {
			return &record, true
		}
	}
	return nil, false
}

// --- Indexed Search (Conceptual) ---
// Index represents a conceptual index (e.g., a map for fast lookup).
type Index map[int]int // Map ID to its index in the original slice

// BuildIndex creates an index for the records based on ID.
func BuildIndex(records []Record) Index {
	index := make(Index)
	for i, record := range records {
		index[record.ID] = i
	}
	return index
}

// IndexedSearchRecords performs a search using the pre-built index.
func IndexedSearchRecords(records []Record, index Index, targetID int) (*Record, bool) {
	if idx, found := index[targetID]; found {
		return &records[idx], true
	}
	return nil, false
}

// --- Sorted Search (Conceptual) ---
// For demonstrating binary search on a sorted slice.
type RecordsByID []Record

func (r RecordsByID) Len() int           { return len(r) }
func (r RecordsByID) Less(i, j int) bool { return r[i].ID < r[j].ID }
func (r RecordsByID) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// BinarySearchRecords performs a binary search on a sorted slice of records.
func BinarySearchRecords(records []Record, targetID int) (*Record, bool) {
	idx := sort.Search(len(records), func(i int) bool {
		return records[i].ID >= targetID
	})
	if idx < len(records) && records[idx].ID == targetID {
		return &records[idx], true
	}
	return nil, false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("--- Conceptual Database Indexing Impact ---")

	const numRecords = 100000 // A large number of records
	records := generateRecords(numRecords)
	targetID := numRecords / 2 // Search for a middle element

	// --- Linear Search ---
	fmt.Println("\n--- Linear Search (Unindexed) ---")
	start := time.Now()
	foundRecord, ok := LinearSearchRecords(records, targetID)
	duration := time.Since(start)
	if ok {
		fmt.Printf("Found record ID %%d (Name: %%s) in %%s\n", foundRecord.ID, foundRecord.Name, duration)
	} else {
		fmt.Printf("Record ID %%d not found in %%s\n", targetID, duration)
	}

	// --- Indexed Search (using a map as a conceptual index) ---
	fmt.Println("\n--- Indexed Search (Conceptual Map Index) ---")
	fmt.Println("Building index...")
	start = time.Now()
	index := BuildIndex(records)
	buildDuration := time.Since(start)
	fmt.Printf("Index built in %%s\n", buildDuration)

	start = time.Now()
	foundRecord, ok = IndexedSearchRecords(records, index, targetID)
	duration = time.Since(start)
	if ok {
		fmt.Printf("Found record ID %%d (Name: %%s) in %%s\n", foundRecord.ID, foundRecord.Name, duration)
	} else {
		fmt.Printf("Record ID %%d not found in %%s\n", targetID, duration)
	}

	// --- Binary Search (on a sorted slice) ---
	fmt.Println("\n--- Binary Search (on Sorted Slice) ---")
	fmt.Println("Sorting records...")
	start = time.Now()
	sort.Sort(RecordsByID(records)) // Sort the records by ID
	sortDuration := time.Since(start)
	fmt.Printf("Records sorted in %%s\n", sortDuration)

	start = time.Now()
	foundRecord, ok = BinarySearchRecords(records, targetID)
	duration = time.Since(start)
	if ok {
		fmt.Printf("Found record ID %%d (Name: %%s) in %%s\n", foundRecord.ID, foundRecord.Name, duration)
	} else {
		fmt.Printf("Record ID %%d not found in %%s\n", targetID, duration)
	}

	fmt.Println("\n--- Database Indexing Example Finished ---")
}
