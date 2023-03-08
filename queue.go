package gblink

import "errors"

type Queue[T any] []T

type QueueError struct {
	error
}

// NewQueue returns a new Queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push pushes the specified value onto the queue.
//
// Push an item onto the queue.
//
// Example:
//
//	q := NewQueue()
//	q.Push(1)
//	q.Push(2)
//	q.Push(3)
//	fmt.Println(q) // [1 2 3]
func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

// Pop removes and returns the first item from the queue.
//
// Pop an item from the queue.
//
// Example:
//
//	q := NewQueue()
//	q.Push(1)
//	q.Push(2)
//	q.Push(3)
//	fmt.Println(q.Pop()) // 1
//	fmt.Println(q.Pop()) // 2
//	fmt.Println(q.Pop()) // 3
func (q *Queue[T]) Pop() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, &QueueError{errors.New("QueueError: queue is empty")}
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, nil
}

// Peek returns the first item from the queue without removing it.
//
// Peek at the first item on the queue.
//
// Example:
//
//	q := NewQueue()
//	q.Push(1)
//	q.Push(2)
//	q.Push(3)
//	fmt.Println(q.Peek()) // 1
//	fmt.Println(q.Peek()) // 1
//	fmt.Println(q.Peek()) // 1
func (q *Queue[T]) Peek() (T, error) {
	if len(*q) == 0 {
		var zero T
		return zero, &QueueError{errors.New("QueueError: queue is empty")}
	}
	return (*q)[0], nil
}

// Len returns the number of items in the queue.
//
// Get the number of items in the queue.
func (q *Queue[T]) Len() int {
	return len(*q)
}

// IsEmpty returns true if the queue is empty.
//
// Check if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
