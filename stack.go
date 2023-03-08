package gblink

import "errors"

type Stack[T any] []T

type StackError struct {
	error
}

// NewStack returns a new Stack.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push pushes the specified value onto the stack.
//
// Push an item onto the stack.
//
// Example:
//
//	s := NewStack()
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	fmt.Println(s) // [1 2 3]
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop removes and returns the top item from the stack.
//
// Pop an item from the stack.
//
// Example:
//
//	s := NewStack()
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	fmt.Println(s.Pop()) // 3
//	fmt.Println(s.Pop()) // 2
//	fmt.Println(s.Pop()) // 1
func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		var zero T
		return zero, &StackError{errors.New("StackError: stack is empty")}
	}
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v, nil
}

// Peek returns the top item from the stack without removing it.
//
// Peek at the top item on the stack.
//
// Example:
//
//	s := NewStack()
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	fmt.Println(s.Peek()) // 3
func (s *Stack[T]) Peek() (T, error) {
	if len(*s) == 0 {
		var zero T
		return zero, &StackError{errors.New("StackError: stack is empty")}
	}
	return (*s)[len(*s)-1], nil
}

// Len returns the number of items in the stack.
//
// Get the number of items in the stack.
//
// Example:
//
//	s := NewStack()
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//	fmt.Println(s.Len()) // 3
func (s *Stack[T]) Len() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty.
//
// Check if the stack is empty.
//
// Example:
//
//		s := NewStack()
//		fmt.Println(s.IsEmpty()) // true
//		s.Push(1)
//		fmt.Println(s.IsEmpty()) // false
//	 s.Pop()
//		fmt.Println(s.IsEmpty()) // true
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
