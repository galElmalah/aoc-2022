package queue

import "fmt"

type Queue[T comparable] struct {
	items []T
}

// Enqueue - Adds an item T to the Q
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue - Removes an element from the Q - FIFO
func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		fmt.Println("Trying to Dequeue from an empty Queue")
	}
	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem
}

// Peek - Look at the head of the Q
func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func (q *Queue[T]) NumberOfItems() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
