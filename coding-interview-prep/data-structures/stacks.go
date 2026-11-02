package main

import (
	"fmt"
	"sync" // For thread-safe stack if needed, though not strictly necessary for basic example
)

// Stack represents a stack data structure
type Stack struct {
	elements []interface{} // Using interface{} to allow any type
	lock     sync.RWMutex  // For concurrent access safety
}

// NewStack creates and returns a new Stack
func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.elements = append(s.elements, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (interface{}, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.IsEmpty() {
		return nil, false
	}

	index := len(s.elements) - 1
	item := s.elements[index]
	s.elements = s.elements[:index] // Remove the element
	return item, true
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if s.IsEmpty() {
		return nil, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.elements)
}

func main() {
	stack := NewStack()
	fmt.Println("Is stack empty?", stack.IsEmpty()) // Expected: true

	stack.Push(10)
	stack.Push("hello")
	stack.Push(true)
	fmt.Println("Stack size after pushes:", stack.Size()) // Expected: 3

	top, ok := stack.Peek()
	if ok {
		fmt.Println("Top element:", top) // Expected: true
	}

	item, ok := stack.Pop()
	if ok {
		fmt.Println("Popped element:", item) // Expected: true
	}
	fmt.Println("Stack size after one pop:", stack.Size()) // Expected: 2

	item, ok = stack.Pop()
	if ok {
		fmt.Println("Popped element:", item) // Expected: hello
	}

	item, ok = stack.Pop()
	if ok {
		fmt.Println("Popped element:", item) // Expected: 10
	}

	item, ok = stack.Pop()
	if !ok {
		fmt.Println("Stack is empty, cannot pop.") // Expected: Stack is empty, cannot pop.
	}

	fmt.Println("Is stack empty?", stack.IsEmpty()) // Expected: true
}
