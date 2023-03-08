package gblink

import "fmt"

type TreeError struct {
	error
}

type TreeNode[K comparable, V any] struct {
	Key      K
	Value    V
	Children []*TreeNode[K, V]
	Parent   *TreeNode[K, V]
}

// NewTreeNode creates a new tree node with the given key and value.
//
// The new node has no children and no parent.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	fmt.Println(node.Key) // "key"
func NewTreeNode[K comparable, V any](key K, value V) *TreeNode[K, V] {
	return &TreeNode[K, V]{Key: key, Value: value, Children: []*TreeNode[K, V]{}, Parent: nil}
}

// AddChild adds the given child to the node's children.
//
// The child's parent is set to the node.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	child := NewTreeNode("child", "value")
//	node.AddChild(child)
//	fmt.Println(node.Children[0].Key) // "child"
//	fmt.Println(child.Parent.Key) // "key"
func (t *TreeNode[K, V]) AddChild(child *TreeNode[K, V]) {
	t.Children = append(t.Children, child)
	child.Parent = t
}

// RemoveChild removes the given child from the node's children.
//
// The child's parent is set to nil.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	child := NewTreeNode("child", "value")
//	node.AddChild(child)
//	node.RemoveChild(child)
//	fmt.Println(node.Children) // []
//	fmt.Println(child.Parent) // nil
func (t *TreeNode[K, V]) RemoveChild(child *TreeNode[K, V]) {
	for i, c := range t.Children {
		if c == child {
			t.Children = append(t.Children[:i], t.Children[i+1:]...)
			child.Parent = nil
			return
		}
	}
}

// RemoveChildKey removes the child with the given key from the node's children.
//
// The child's parent is set to nil.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	child := NewTreeNode("child", "value")
//	node.AddChild(child)
//	node.RemoveChildKey("child")
//	fmt.Println(node.Children) // []
//	fmt.Println(child.Parent) // nil
func (t *TreeNode[K, V]) RemoveChildKey(key K) {
	for i, c := range t.Children {
		if c.Key == key {
			t.Children = append(t.Children[:i], t.Children[i+1:]...)
			c.Parent = nil
			return
		}
	}
}

// RemoveChildAt removes the child at the given index from the node's children.
//
// The child's parent is set to nil.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	child := NewTreeNode("child", "value")
//	node.AddChild(child)
//	node.RemoveChildAt(0)
//	fmt.Println(node.Children) // []
//	fmt.Println(child.Parent) // nil
func (t *TreeNode[K, V]) RemoveChildAt(index int) error {
	if index < 0 || index >= len(t.Children) {
		return &TreeError{fmt.Errorf("TreeError: index out of range")}
	}
	t.Children[index].Parent = nil
	t.Children = append(t.Children[:index], t.Children[index+1:]...)
	return nil
}

// LenChildren returns the number of children the node has.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	child := NewTreeNode("child", "value")
//	node.AddChild(child)
//	fmt.Println(node.LenChildren()) // 1
func (t *TreeNode[K, V]) LenChildren() int {
	return len(t.Children)
}

// SetParent sets the node's parent to the given parent.
//
// The parent's children are updated to include the node.
//
// Example:
//
//	node := NewTreeNode("key", "value")
//	parent := NewTreeNode("parent", "value")
//	node.SetParent(parent)
//	fmt.Println(node.Parent.Key) // "parent"
//	fmt.Println(parent.Children[0].Key) // "key"
func (t *TreeNode[K, V]) SetParent(parent *TreeNode[K, V]) {
	t.Parent = parent
	parent.Children = append(parent.Children, t)
}
