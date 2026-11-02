package main

import (
	"fmt"
	"sync"
	"time"
)

// Node represents a replica in a distributed system.
type Node struct {
	ID    int
	Data  map[string]string
	mu    sync.RWMutex
	delay time.Duration // Simulate network delay for propagation
}

// NewNode creates a new Node.
func NewNode(id int, delay time.Duration) *Node {
	return &Node{
		ID:    id,
		Data:  make(map[string]string),
		delay: delay,
	}
}

// Write simulates a write operation to the node.
func (n *Node) Write(key, value string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Data[key] = value
	fmt.Printf("Node %d: Wrote '%s' = '%s'\n", n.ID, key, value)
}

// Read simulates a read operation from the node.
func (n *Node) Read(key string) (string, bool) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	value, ok := n.Data[key]
	fmt.Printf("Node %d: Read '%s' = '%s' (found: %t)\n", n.ID, key, value, ok)
	return value, ok
}

// Propagate simulates data propagation to other nodes.
func (n *Node) Propagate(key, value string, otherNodes []*Node) {
	time.Sleep(n.delay) // Simulate network delay
	for _, other := range otherNodes {
		if other.ID != n.ID {
			go func(targetNode *Node) {
				time.Sleep(targetNode.delay) // Simulate delay to target
				targetNode.mu.Lock()
				targetNode.Data[key] = value
				targetNode.mu.Unlock()
				fmt.Printf("Node %d: Propagated '%s' = '%s' to Node %d\n", n.ID, key, value, targetNode.ID)
			}(other)
		}
	}
}

func main() {
	fmt.Println("--- Conceptual Eventual Consistency Simulation ---")

	// Create multiple nodes
	node1 := NewNode(1, 100*time.Millisecond)
	node2 := NewNode(2, 200*time.Millisecond)
	node3 := NewNode(3, 50*time.Millisecond)
	nodes := []*Node{node1, node2, node3}

	var wg sync.WaitGroup

	// Client 1 writes to Node 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		node1.Write("user_name", "Alice")
		node1.Propagate("user_name", "Alice", nodes)
	}()

	// Client 2 reads from Node 2 shortly after
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(150 * time.Millisecond) // Read before full propagation
		node2.Read("user_name")
	}()

	// Client 3 reads from Node 3 even later
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond) // Read after more propagation
		node3.Read("user_name")
	}()

	// Client 1 updates to Node 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		node1.Write("user_name", "Alicia")
		node1.Propagate("user_name", "Alicia", nodes)
	}()

	// Client 2 reads from Node 1 after update
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(600 * time.Millisecond)
		node1.Read("user_name")
	}()

	// Client 3 reads from Node 2 after update
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(800 * time.Millisecond)
		node2.Read("user_name")
	}()

	wg.Wait()
	fmt.Println("\n--- All operations completed. Waiting for final consistency... ---")
	time.Sleep(500 * time.Millisecond) // Give time for all propagations to finish

	fmt.Println("\n--- Final State ---")
	node1.Read("user_name")
	node2.Read("user_name")
	node3.Read("user_name")

	fmt.Println("\n--- Consistency Model Example Finished ---")
}
