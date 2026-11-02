package main

import "fmt"

// Node represents a node in the singly linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the singly linked list
type LinkedList struct {
	head *Node
	length int
}

// Append adds a new node with the given data to the end of the list
func (l *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if l.head == nil {
		l.head = newNode
		l.length = 1
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	l.length++
}

// Prepend adds a new node with the given data to the beginning of the list
func (l *LinkedList) Prepend(data int) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
	l.length++
}

// InsertAt inserts a new node with the given data at a specific position
func (l *LinkedList) InsertAt(data, position int) error {
	if position < 0 || position > l.length {
		return fmt.Errorf("position out of bounds")
	}
	if position == 0 {
		l.Prepend(data)
		return nil
	}
	if position == l.length {
		l.Append(data)
		return nil
	}

	newNode := &Node{data: data}
	current := l.head
	for i := 0; i < position-1; i++ {
		current = current.next
	}
	newNode.next = current.next
	current.next = newNode
	l.length++
	return nil
}

// Delete removes the first node with the given data
func (l *LinkedList) Delete(data int) error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	if l.head.data == data {
		l.head = l.head.next
		l.length--
		return nil
	}

	current := l.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}

	if current.next == nil {
		return fmt.Errorf("data not found in list")
	}

	current.next = current.next.next
	l.length--
	return nil
}

// Search finds a node with the given data and returns true if found, false otherwise
func (l *LinkedList) Search(data int) bool {
	current := l.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

// Display prints the elements of the list
func (l *LinkedList) Display() {
	current := l.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	fmt.Println("Initial list:")
	list.Display()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("After appending 10, 20, 30:")
	list.Display() // Expected: 10 -> 20 -> 30 -> nil

	list.Prepend(5)
	fmt.Println("After prepending 5:")
	list.Display() // Expected: 5 -> 10 -> 20 -> 30 -> nil

	list.InsertAt(15, 2)
	fmt.Println("After inserting 15 at position 2:")
	list.Display() // Expected: 5 -> 10 -> 15 -> 20 -> 30 -> nil

	fmt.Printf("Search for 20: %v\n", list.Search(20)) // Expected: true
	fmt.Printf("Search for 100: %v\n", list.Search(100)) // Expected: false

	list.Delete(15)
	fmt.Println("After deleting 15:")
	list.Display() // Expected: 5 -> 10 -> 20 -> 30 -> nil

	list.Delete(5)
	fmt.Println("After deleting 5 (head):")
	list.Display() // Expected: 10 -> 20 -> 30 -> nil

	list.Delete(30)
	fmt.Println("After deleting 30 (tail):")
	list.Display() // Expected: 10 -> 20 -> nil

	err := list.Delete(100)
	if err != nil {
		fmt.Println("Error deleting 100:", err) // Expected: Error deleting 100: data not found in list
	}

	fmt.Printf("List length: %d\n", list.length) // Expected: 2
}
