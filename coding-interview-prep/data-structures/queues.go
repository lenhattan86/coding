package main

import (
	"fmt"
	"sync"
)

// Queue represents a queue data structure
type Queue struct {
	elements []interface{}
	lock     sync.RWMutex
}

// NewQueue creates and returns a new Queue
func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
	}
}

// Enqueue adds an element to the back of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.elements = append(q.elements, item)
}

// Dequeue removes and returns the front element from the queue
func (q *Queue) Dequeue() (interface{}, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.IsEmpty() {
		return nil, false
	}

	item := q.elements[0]
	q.elements = q.elements[1:] // Remove the first element
	return item, true
}

// Peek returns the front element of the queue without removing it
func (q *Queue) Peek() (interface{}, bool) {
	q.lock.RLock()
	defer q.lock.RUnlock()

	if q.IsEmpty() {
		return nil, false
	}
	return q.elements[0], true
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.elements) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.elements)
}

func main() {
	queue := NewQueue()
	fmt.Println("Is queue empty?", queue.IsEmpty()) // Expected: true

	queue.Enqueue(10)
	queue.Enqueue("hello")
	queue.Enqueue(true)
	fmt.Println("Queue size after enqueues:", queue.Size()) // Expected: 3

	front, ok := queue.Peek()
	if ok {
		fmt.Println("Front element:", front) // Expected: 10
	}

	item, ok := queue.Dequeue()
	if ok {
		fmt.Println("Dequeued element:", item) // Expected: 10
	}
	fmt.Println("Queue size after one dequeue:", queue.Size()) // Expected: 2

	item, ok = queue.Dequeue()
	if ok {
		fmt.Println("Dequeued element:", item) // Expected: hello
	}

	item, ok = queue.Dequeue()
	if ok {
		fmt.Println("Dequeued element:", item) // Expected: true
	}

	item, ok = queue.Dequeue()
	if !ok {
		fmt.Println("Queue is empty, cannot dequeue.") // Expected: Queue is empty, cannot dequeue.
	}

	fmt.Println("Is queue empty?", queue.IsEmpty()) // Expected: true
}
