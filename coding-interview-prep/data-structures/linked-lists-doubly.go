package main

import "fmt"

// DoublyNode represents a node in the doubly linked list
type DoublyNode struct {
	data int
	next *DoublyNode
	prev *DoublyNode
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int
}

// Append adds a new node with the given data to the end of the list
func (l *DoublyLinkedList) Append(data int) {
	newNode := &DoublyNode{data: data, next: nil, prev: nil}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
	l.length++
}

// Prepend adds a new node with the given data to the beginning of the list
func (l *DoublyLinkedList) Prepend(data int) {
	newNode := &DoublyNode{data: data, next: l.head, prev: nil}
	if l.head == nil {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}
	l.head = newNode
	l.length++
}

// InsertAfter inserts a new node with the given data after a specific node
func (l *DoublyLinkedList) InsertAfter(nodeData, newData int) error {
	current := l.head
	for current != nil && current.data != nodeData {
		current = current.next
	}

	if current == nil {
		return fmt.Errorf("node with data %d not found", nodeData)
	}

	if current == l.tail { // Insert at the end
		l.Append(newData)
	} else {
		newNode := &DoublyNode{data: newData, next: current.next, prev: current}
		current.next.prev = newNode
		current.next = newNode
		l.length++
	}
	return nil
}

// Delete removes the first node with the given data
func (l *DoublyLinkedList) Delete(data int) error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}

	current := l.head
	for current != nil && current.data != data {
		current = current.next
	}

	if current == nil {
		return fmt.Errorf("data %d not found in list", data)
	}

	if current == l.head { // Deleting the head
		l.head = current.next
		if l.head != nil {
			l.head.prev = nil
		} else {
			l.tail = nil // List became empty
		}
	} else if current == l.tail { // Deleting the tail
		l.tail = current.prev
		l.tail.next = nil
	} else { // Deleting a middle node
		current.prev.next = current.next
		current.next.prev = current.prev
	}
	l.length--
	return nil
}

// DisplayForward prints the elements of the list from head to tail
func (l *DoublyLinkedList) DisplayForward() {
	current := l.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Print("Forward: ")
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// DisplayBackward prints the elements of the list from tail to head
func (l *DoublyLinkedList) DisplayBackward() {
	current := l.tail
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	fmt.Print("Backward: ")
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func main() {
	list := DoublyLinkedList{}
	fmt.Println("Initial list:")
	list.DisplayForward()

	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("After appending 10, 20, 30:")
	list.DisplayForward() // Expected: 10 <-> 20 <-> 30 <-> nil
	list.DisplayBackward() // Expected: 30 <-> 20 <-> 10 <-> nil

	list.Prepend(5)
	fmt.Println("After prepending 5:")
	list.DisplayForward() // Expected: 5 <-> 10 <-> 20 <-> 30 <-> nil
	list.DisplayBackward() // Expected: 30 <-> 20 <-> 10 <-> 5 <-> nil

	list.InsertAfter(10, 15)
	fmt.Println("After inserting 15 after 10:")
	list.DisplayForward() // Expected: 5 <-> 10 <-> 15 <-> 20 <-> 30 <-> nil
	list.DisplayBackward() // Expected: 30 <-> 20 <-> 15 <-> 10 <-> 5 <-> nil

	list.Delete(20)
	fmt.Println("After deleting 20:")
	list.DisplayForward() // Expected: 5 <-> 10 <-> 15 <-> 30 <-> nil
	list.DisplayBackward() // Expected: 30 <-> 15 <-> 10 <-> 5 <-> nil

	list.Delete(5) // Delete head
	fmt.Println("After deleting 5 (head):")
	list.DisplayForward() // Expected: 10 <-> 15 <-> 30 <-> nil
	list.DisplayBackward() // Expected: 30 <-> 15 <-> 10 <-> nil

	list.Delete(30) // Delete tail
	fmt.Println("After deleting 30 (tail):")
	list.DisplayForward() // Expected: 10 <-> 15 <-> nil
	list.DisplayBackward() // Expected: 15 <-> 10 <-> nil

	err := list.Delete(100)
	if err != nil {
		fmt.Println("Error deleting 100:", err) // Expected: data 100 not found in list
	}

	fmt.Printf("List length: %d\n", list.length) // Expected: 2
}
