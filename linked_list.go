package gblink

import "fmt"

type LikedListNode[T comparable] struct {
	Value T
	Next  *LikedListNode[T]
}

type LikedList[T comparable] struct {
	Head *LikedListNode[T]
	Tail *LikedListNode[T]
}

type LikedListError struct {
	error
}

func NewLikedList[T comparable]() *LikedList[T] {
	return &LikedList[T]{}
}

// Len returns the number of elements in the list.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Len() // 3
func (l *LikedList[T]) Len() int {
	count := 0
	for node := l.Head; node != nil; node = node.Next {
		count++
	}
	return count
}

// Append adds a new element with the given value to the end of the list.
//
// The complexity is O(1).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Len() // 3
func (l *LikedList[T]) Append(value T) {
	node := &LikedListNode[T]{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	l.Tail = node
}

// Prepend adds a new element with the given value to the beginning of the list.
//
// The complexity is O(1).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Prepend(1)
//	list.Prepend(2)
//	list.Prepend(3)
//	list.Len() // 3
//	fmt.Println(list.Head.Value) // 3
func (l *LikedList[T]) Prepend(value T) {
	node := &LikedListNode[T]{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	node.Next = l.Head
	l.Head = node
}

// Insert adds a new element with the given value after the n-th element of the list.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Insert(2, 4)
//	list.Len() // 4
//	fmt.Println(list.Head.Next.Next.Value) // 4
func (l *LikedList[T]) Insert(n int, value T) error {
	if n < 0 {
		return &LikedListError{fmt.Errorf("LikedListError: index out of range")}
	}
	if n == 0 {
		l.Prepend(value)
		return nil
	}
	if n == l.Len() {
		l.Append(value)
		return nil
	}
	node := l.Head
	for i := 0; i < n-1; i++ {
		node = node.Next
	}
	newNode := &LikedListNode[T]{Value: value}
	newNode.Next = node.Next
	node.Next = newNode
	return nil
}

// Remove removes the n-th element of the list and returns its value.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Remove(1) // 2
//	list.Len() // 2
func (l *LikedList[T]) Remove(n int) (T, error) {
	if n < 0 || n >= l.Len() {
		var zero T
		return zero, &LikedListError{fmt.Errorf("LikedListError: index out of range")}
	}
	if n == 0 {
		value := l.Head.Value
		l.Head = l.Head.Next
		return value, nil
	}
	node := l.Head
	for i := 0; i < n-1; i++ {
		node = node.Next
	}
	value := node.Next.Value
	node.Next = node.Next.Next
	return value, nil
}

// Get returns the value of the n-th element of the list.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Get(1) // 2
func (l *LikedList[T]) Get(n int) (T, error) {
	if n < 0 || n >= l.Len() {
		var zero T
		return zero, &LikedListError{fmt.Errorf("LikedListError: index out of range")}
	}
	node := l.Head
	for i := 0; i < n; i++ {
		node = node.Next
	}
	return node.Value, nil
}

// IndexOf returns the index of the first element with the given value.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.IndexOf(2) // 1
func (l *LikedList[T]) IndexOf(value T) int {
	node := l.Head
	for i := 0; i < l.Len(); i++ {
		if node.Value == value {
			return i
		}
		node = node.Next
	}
	return -1
}

// Contains returns true if the list contains an element with the given value.
//
// The complexity is O(n).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Contains(2) // true
func (l *LikedList[T]) Contains(value T) bool {
	return l.IndexOf(value) != -1
}

// Clear removes all elements from the list.
//
// The complexity is O(1).
//
// Example:
//
//	list := NewLikedList[int]()
//	list.Append(1)
//	list.Append(2)
//	list.Append(3)
//	list.Clear()
//	list.Len() // 0
func (l *LikedList[T]) Clear() {
	l.Head = nil
	l.Tail = nil
}
