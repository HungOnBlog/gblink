package gblink

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeError struct {
	error
}

// TreeNode is a node in a tree.
//
// The zero value for TreeNode is an empty tree node ready to use.
type TreeNode[K constraints.Ordered, V any] struct {
	Key   K
	Value V
	Left  *TreeNode[K, V]
	Right *TreeNode[K, V]
}

// Tree is a tree implementation.
//
// The zero value for Tree is an empty tree ready to use.
type Tree[K constraints.Ordered, V any] struct {
	Root *TreeNode[K, V]
}

// NewTree returns a new Tree.
func NewTree[K constraints.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{
		Root: nil,
	}
}

// Set sets the value for the given key.
//
// The complexity is O(log n).
//
// Example:
//
//		tree := NewTree[int, string]()
//		tree.Set(1, "one")
//		tree.Set(2, "two")
//		tree.Set(3, "three")
//		tree.Set(4, "four")
//		tree.Set(5, "five")
//	 fmt.Println(tree.Root.Value) // five
func (t *Tree[K, V]) Set(key K, value V) {
	t.Root = t.set(t.Root, key, value)
}

func (t *Tree[K, V]) set(node *TreeNode[K, V], key K, value V) *TreeNode[K, V] {
	if node == nil {
		return &TreeNode[K, V]{
			Key:   key,
			Value: value,
			Left:  nil,
			Right: nil,
		}
	}
	if key < node.Key {
		node.Left = t.set(node.Left, key, value)
	} else if key > node.Key {
		node.Right = t.set(node.Right, key, value)
	} else {
		node.Value = value
	}
	return node
}

// Get returns the value for the given key.
//
// The complexity is O(log n).
//
// Example:
//
//	tree := NewTree[int, string]()
//	tree.Set(1, "one")
//	tree.Set(2, "two")
//	tree.Set(3, "three")
//	tree.Set(4, "four")
//	tree.Set(5, "five")
//	fmt.Println(tree.Get(3)) // three
func (t *Tree[K, V]) Get(key K) (V, error) {
	return t.get(t.Root, key)
}

func (t *Tree[K, V]) get(node *TreeNode[K, V], key K) (V, error) {
	if node == nil {
		var zero V
		return zero, &TreeError{fmt.Errorf("TreeError: key not found: %v", key)}
	}
	if key < node.Key {
		return t.get(node.Left, key)
	} else if key > node.Key {
		return t.get(node.Right, key)
	}
	return node.Value, nil
}

// Len returns the number of elements in the tree.
//
// The complexity is O(n).
//
// Example:
//
//	tree := NewTree[int, string]()
//	tree.Set(1, "one")
//	tree.Set(2, "two")
//	tree.Set(3, "three")
//	tree.Set(4, "four")
//	tree.Set(5, "five")
//	fmt.Println(tree.Len()) // 5
func (t *Tree[K, V]) Len() int {
	return t.len(t.Root)
}

func (t *Tree[K, V]) len(node *TreeNode[K, V]) int {
	if node == nil {
		return 0
	}
	return 1 + t.len(node.Left) + t.len(node.Right)
}

// Delete deletes the value for the given key.
//
// The complexity is O(log n).
//
// Example:
//
//	tree := NewTree[int, string]()
//	tree.Set(1, "one")
//	tree.Set(2, "two")
//	tree.Set(3, "three")
//	tree.Set(4, "four")
//	tree.Set(5, "five")
//	tree.Delete(3)
//	fmt.Println(tree.Get(3)) // TreeError: key not found: 3
func (t *Tree[K, V]) Delete(key K) {
	t.Root = t.delete(t.Root, key)
}

func (t *Tree[K, V]) DeleteMin(node *TreeNode[K, V]) *TreeNode[K, V] {
	if node.Left == nil {
		return node.Right
	}
	node.Left = t.DeleteMin(node.Left)
	return node
}

func (t *Tree[K, V]) Min(node *TreeNode[K, V]) *TreeNode[K, V] {
	if node.Left == nil {
		return node
	}
	return t.Min(node.Left)
}

func (t *Tree[K, V]) delete(node *TreeNode[K, V], key K) *TreeNode[K, V] {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = t.delete(node.Left, key)
	} else if key > node.Key {
		node.Right = t.delete(node.Right, key)
	} else {
		if node.Right == nil {
			return node.Left
		}
		if node.Left == nil {
			return node.Right
		}
		temp := node
		node = t.Min(temp.Right)
		node.Right = t.DeleteMin(temp.Right)
		node.Left = temp.Left
	}
	return node
}

// Keys returns a slice of keys in the tree.
//
// The complexity is O(n).
//
// Example:
//
//	tree := NewTree[int, string]()
//	tree.Set(1, "one")
//	tree.Set(2, "two")
//	tree.Set(3, "three")
//	tree.Set(4, "four")
//	tree.Set(5, "five")
//	fmt.Println(tree.Keys()) // [1 2 3 4 5]
func (t *Tree[K, V]) Keys() []K {
	keys := make([]K, 0, t.Len())
	t.keys(t.Root, &keys)
	return keys
}

func (t *Tree[K, V]) keys(node *TreeNode[K, V], keys *[]K) {
	if node == nil {
		return
	}
	t.keys(node.Left, keys)
	*keys = append(*keys, node.Key)
	t.keys(node.Right, keys)
}

// Max returns the maximum key in the tree.
//
// The complexity is O(log n).
//
// Example:
//
//	tree := NewTree[int, string]()
//	tree.Set(1, "one")
//	tree.Set(2, "two")
//	tree.Set(3, "three")
//	tree.Set(4, "four")
//	tree.Set(5, "five")
//	fmt.Println(tree.Max()) // 5
func (t *Tree[K, V]) Max() (K, error) {
	return t.max(t.Root)
}

func (t *Tree[K, V]) max(node *TreeNode[K, V]) (K, error) {
	if node == nil {
		var zero K
		return zero, &TreeError{fmt.Errorf("TreeError: tree is empty")}
	}
	if node.Right == nil {
		return node.Key, nil
	}
	return t.max(node.Right)
}
