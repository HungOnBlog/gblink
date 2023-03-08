package gblink

import (
	"fmt"
	"hash"
)

type HashTableError struct {
	error
}

// HashTable is a hash table implementation.
//
// The zero value for HashTable is an empty hash table ready to use.
//
// The HashTable type is not safe for concurrent use by multiple goroutines without.
type HashTable[K comparable, V comparable] struct {
	Hasher hash.Hash64
	Table  map[uint64]*LikedList[V]
}

// NewHashTable returns a new HashTable.
func NewHashTable[K comparable, V comparable](hasher hash.Hash64) *HashTable[K, V] {
	return &HashTable[K, V]{
		Table:  make(map[uint64]*LikedList[V]),
		Hasher: hasher,
	}
}

// Set sets the value for the given key.
//
// The complexity is O(1).
//
// Example:
//
//	table := NewHashTable[int, string]()
//	table.Set(1, "one")
//	table.Set(2, "two")
//	table.Set(3, "three")
//	table.Set(4, "four")
//	table.Set(5, "five")
//	table.Set(6, "six")
func (t *HashTable[K, V]) Set(key K, value V) {
	// Hash the key.
	t.Hasher.Reset()
	t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
	hash := t.Hasher.Sum64()
	if _, ok := t.Table[hash]; !ok {
		t.Table[hash] = NewLikedList[V]()
	}
	t.Table[hash].Append(value)
}

// Get returns the value for the given key.
//
// The complexity is O(1).
//
// Example:
//
//		table := NewHashTable[int, string]()
//		table.Set(1, "one")
//		table.Set(2, "two")
//		table.Set(3, "three")
//	 v , err := table.Get(2)
//	 if err != nil {
//	     panic(err)
//	 }
//	 fmt.Println(v) // two
func (t *HashTable[K, V]) Get(key K) (V, error) {
	// Hash the key.
	t.Hasher.Reset()
	t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
	hash := t.Hasher.Sum64()
	if _, ok := t.Table[hash]; !ok {
		var zero V
		return zero, &HashTableError{error: fmt.Errorf("HashTableError: key not found")}
	}
	return t.Table[hash].Head.Value, nil
}

// Len returns the number of elements in the hash table.
//
// The complexity is O(n).
//
// Example:
//
//		table := NewHashTable[int, string]()
//		table.Set(1, "one")
//		table.Set(2, "two")
//		table.Set(3, "three")
//		table.Set(4, "four")
//		table.Set(5, "five")
//	    fmt.Println(table.Len()) // 5
func (t *HashTable[K, V]) Len() int {
	count := 0
	for _, list := range t.Table {
		count += list.Len()
	}
	return count
}

// Clear removes all elements from the hash table.
//
// The complexity is O(n).
//
// Example:
//
//	table := NewHashTable[int, string]()
//	table.Set(1, "one")
//	table.Set(2, "two")
//	table.Set(3, "three")
//	table.Clear()
//	fmt.Println(table.Len()) // 0
func (t *HashTable[K, V]) Clear() {
	t.Table = make(map[uint64]*LikedList[V])
}

// Delete removes the element with the given key from the hash table.
//
// NOTE!: This is dangerous, because it will remove all elements with the same hash.
//
// The complexity is O(1).
//
// Example:
//
//	table := NewHashTable[int, string]()
//	table.Set(1, "one")
//	table.Set(2, "two")
//	table.Set(3, "three")
//	table.Delete(2)
//	fmt.Println(table.Len()) // 2
func (t *HashTable[K, V]) Delete(key K) {
	// Hash the key.
	t.Hasher.Reset()
	t.Hasher.Write([]byte(fmt.Sprintf("%v", key)))
	hash := t.Hasher.Sum64()
	delete(t.Table, hash)
}
