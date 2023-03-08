package gblink

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ArrayError struct {
	error
}

type Array[T constraints.Ordered] []T

// Len returns the length of the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Len()) // 3
func (a *Array[T]) Len() int {
	return len(*a)
}

// Append appends the given value to the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array) // [1 2 3]
func (a *Array[T]) Append(value T) {
	*a = append(*a, value)
}

// AppendAll appends the given values to the array.
//
// Example:
//
//	var array Array[int]
//	array.AppendAll(1, 2, 3)
//	fmt.Println(array) // [1 2 3]
func (a *Array[T]) AppendAll(values ...T) {
	*a = append(*a, values...)
}

// AppendArray appends the given array to the array.
//
// Example:
//
//	var array1 Array[int]
//	array1.AppendArray(&Array[int]{1, 2, 3})
//	fmt.Println(array1) // [1 2 3]
func (a *Array[T]) AppendArray(array *Array[T]) {
	*a = append(*a, *array...)
}

// AppendArrayAll appends the given arrays to the array.
//
// Example:
//
//	var array1 Array[int]
//	array1.AppendArrayAll(&Array[int]{1, 2, 3}, &Array[int]{4, 5, 6})
//	fmt.Println(array1) // [1 2 3 4 5 6]
func (a *Array[T]) AppendArrayAll(arrays ...*Array[T]) {
	for _, array := range arrays {
		*a = append(*a, *array...)
	}
}

// Clear clears the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
func (a *Array[T]) Clear() {
	*a = (*a)[:0]
}

// Contains returns true if the array contains the given value.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Contains(2)) // true
func (a *Array[T]) Contains(value T) bool {
	for _, v := range *a {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the array contains all the given values.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.ContainsAll(1, 2)) // true
func (a *Array[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !a.Contains(value) {
			return false
		}
	}
	return true
}

// Some returns true if the array contains at least one value for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Some(func(value int) bool {
//		return value > 2
//	})) // true
func (a *Array[T]) Some(predicate func(T) bool) bool {
	for _, v := range *a {
		if predicate(v) {
			return true
		}
	}
	return false
}

// Every returns true if the array contains all values for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Every(func(value int) bool {
//		return value > 0
//	})) // true
func (a *Array[T]) Every(predicate func(T) bool) bool {
	for _, v := range *a {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// Filter returns a new array containing all values for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Filter(func(value int) bool {
//		return value > 1
//	})) // [2 3]
func (a *Array[T]) Filter(predicate func(T) bool) *Array[T] {
	var array Array[T]
	for _, v := range *a {
		if predicate(v) {
			array.Append(v)
		}
	}
	return &array
}

// Find returns the first value for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Find(func(value int) bool {
//		return value > 1
//	})) // 2
func (a *Array[T]) Find(predicate func(T) bool) T {
	for _, v := range *a {
		if predicate(v) {
			return v
		}
	}
	var zero T
	return zero
}

// FindIndex returns the index of the first value for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.FindIndex(func(value int) bool {
//		return value > 1
//	})) // 1
func (a *Array[T]) FindIndex(predicate func(T) bool) int {
	for i, v := range *a {
		if predicate(v) {
			return i
		}
	}
	return -1
}

// FindLast returns the last value for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.FindLast(func(value int) bool {
//		return value > 1
//	})) // 3
func (a *Array[T]) FindLast(predicate func(T) bool) T {
	for i := len(*a) - 1; i >= 0; i-- {
		if predicate((*a)[i]) {
			return (*a)[i]
		}
	}
	var zero T
	return zero
}

// FindLastIndex returns the index of the last value for which the given predicate returns true.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.FindLastIndex(func(value int) bool {
//		return value > 1
//	})) // 2
func (a *Array[T]) FindLastIndex(predicate func(T) bool) int {
	for i := len(*a) - 1; i >= 0; i-- {
		if predicate((*a)[i]) {
			return i
		}
	}
	return -1
}

// Each calls the given function for each (index, value) pair in the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.Each(func(index int, value int) {
//		fmt.Println(index, value)
//	})
func (a *Array[T]) Each(f func(int, T)) {
	for i, v := range *a {
		f(i, v)
	}
}

// EachIndex calls the given function for each index in the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.EachIndex(func(index int) {
//		fmt.Println(index)
//	})
func (a *Array[T]) EachIndex(f func(int)) {
	for i := range *a {
		f(i)
	}
}

// EachValue calls the given function for each value in the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.EachValue(func(value int) {
//		fmt.Println(value)
//	})
func (a *Array[T]) EachValue(f func(T)) {
	for _, v := range *a {
		f(v)
	}
}

// Map returns a new array containing the values returned by the given function for each value in the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.Map(func(value int) int {
//		return value + 1
//	})) // [2 3 4]
func (a *Array[T]) Map(f func(T) T) *Array[T] {
	var array Array[T]
	for _, v := range *a {
		array.Append(f(v))
	}
	return &array
}

// SortBy sorts the array using the sort function (which should be a function that takes two
// values and returns -1, 0, or 1).
// -1 means the first value is less than the second value
// 0 means the first value is equal to the second value
// 1 means the first value is greater than the second value
//
// Example:
//
//	var array Array[int]
//	array.Append(3)
//	array.Append(1)
//	array.Append(2)
//	array.SortBy(func(a int, b int) int {
//		if a < b {
//			return -1
//		} else if a > b {
//			return 1
//		} else {
//			return 0
//		}
//	})
//	fmt.Println(array) // [1 2 3]
func (a *Array[T]) SortBy(compare func(T, T) int) {
	// Using the compare function with QuickSort
	quickSort(a, 0, len(*a)-1, compare)
}

// quickSort is a helper function for SortBy
func quickSort[T constraints.Ordered](a *Array[T], lo int, hi int, compare func(T, T) int) {
	if lo < hi {
		p := partition(a, lo, hi, compare)
		quickSort(a, lo, p-1, compare)
		quickSort(a, p+1, hi, compare)
	}
}

// partition is a helper function for quickSort
func partition[T constraints.Ordered](a *Array[T], lo int, hi int, compare func(T, T) int) int {
	pivot := (*a)[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if compare((*a)[j], pivot) < 0 {
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
			i++
		}
	}
	(*a)[i], (*a)[hi] = (*a)[hi], (*a)[i]
	return i
}

func compareAsc[T constraints.Ordered](a T, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func compareDes[T constraints.Ordered](a T, b T) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	} else {
		return 0
	}
}

// Sort sorts the array (ascending or descending), default is ascending.
//
// Example:
//
//	var array Array[int]
//	array.Append(3)
//	array.Append(1)
//	array.Append(2)
//	array.Sort()
//	fmt.Println(array) // [1 2 3]
func (a *Array[T]) Sort(isAscending bool) {
	// Using QuickSort
	if isAscending {
		quickSort(a, 0, len(*a)-1, compareAsc[T])
	} else {
		quickSort(a, 0, len(*a)-1, compareDes[T])
	}
}

// Reverse reverses the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.Reverse()
//	fmt.Println(array) // [3 2 1]
func (a *Array[T]) Reverse() {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	}
}

// Copy returns a copy of the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	arrayCopy := array.Copy()
//	fmt.Println(arrayCopy) // [1 2 3]
func (a *Array[T]) Copy() *Array[T] {
	var array Array[T]
	for _, v := range *a {
		array.Append(v)
	}
	return &array
}

// IndexOf returns the index of the first occurrence of the given value in the array.
// If the value is not found, -1 is returned.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	fmt.Println(array.IndexOf(2)) // 1
//	fmt.Println(array.IndexOf(4)) // -1
func (a *Array[T]) IndexOf(value T) int {
	for i, v := range *a {
		if v == value {
			return i
		}
	}
	return -1
}

// Insert inserts the given value at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.Insert(1, 4)
//	fmt.Println(array) // [1 4 2 3]
func (a *Array[T]) Insert(index int, value T) error {
	if index < 0 || index > len(*a) {
		return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
	}
	*a = append(*a, value)
	copy((*a)[index+1:], (*a)[index:])
	(*a)[index] = value
	return nil
}

// Remove removes the value at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.Remove(1)
//	fmt.Println(array) // [1 3]
func (a *Array[T]) Remove(index int) error {
	if index < 0 || index >= len(*a) {
		return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
	}
	*a = append((*a)[:index], (*a)[index+1:]...)
	return nil
}

// RemoveRange removes the values in the range of start and end.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	array.RemoveRange(1, 2)
//	fmt.Println(array) // [1 3]
func (a *Array[T]) RemoveRange(start int, end int) {
	*a = append((*a)[:start], (*a)[end:]...)
}

// Slice returns a slice of the array.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	slice := array.Slice(1, 2)
//	fmt.Println(slice) // [2]
func (a *Array[T]) Slice(start int, end int) Array[T] {
	var array Array[T]
	for i := start; i < end; i++ {
		array.Append((*a)[i])
	}
	return array
}

// InsertArray inserts the given array at the given index.
// If the index is out of bounds, an error is returned.
//
// Example:
//
//	var array1 Array[int]
//	array1.Append(1)
//	array1.Append(2)
//	array1.Append(3)
//	var array2 Array[int]
//	array2.Append(4)
//	array2.Append(5)
//	array2.Append(6)
//	array1.InsertArray(1, array2)
//	fmt.Println(array1) // [1 4 5 6 2 3]
func (a *Array[T]) InsertArray(index int, array Array[T]) error {
	if index < 0 || index > len(*a) {
		return ArrayError{fmt.Errorf("ArrayError: %d index out of range", index)}
	}
	*a = append(*a, array...)
	copy((*a)[index+len(array):], (*a)[index:])
	copy((*a)[index:], array)
	return nil
}

// Reduce applies a function against an accumulator and each value of the array (from left-to-right) as to reduce it to a single value.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	sum := array.Reduce(func(accumulator, value int) int {
//		return accumulator + value
//	}, 0)
//	fmt.Println(sum) // 6
func (a *Array[T]) Reduce(fn func(accumulator, value T) T, accumulator T) T {
	for _, v := range *a {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

// ReduceIf applies a function against an accumulator and each value of the array (from left-to-right) as to reduce it to a single value.
// The function is only applied to the values that satisfy the given predicate.
//
// Example:
//
//	var array Array[int]
//	array.Append(1)
//	array.Append(2)
//	array.Append(3)
//	sum := array.ReduceIf(func(accumulator, value int) int {
//		return accumulator + value
//	}, 0, func(value int) bool {
//		return value > 1
//	})
//	fmt.Println(sum) // 5
func (a *Array[T]) ReduceIf(fn func(accumulator, value T) T, predicate func(value T) bool, accumulator T) T {
	for _, v := range *a {
		if predicate(v) {
			accumulator = fn(accumulator, v)
		}
	}
	return accumulator
}
