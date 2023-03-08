package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_AddChild(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	node.AddChild(child)

	assert.Equal(node.Children[0].Key, "child")
}

func TestTree_RemoveChild(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	node.AddChild(child)
	node.RemoveChild(child)

	assert.Equal(0, len(node.Children))
}

func TestTree_RemoveChildKey(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	node.AddChild(child)
	node.RemoveChildKey("child")

	assert.Equal(0, len(node.Children))
}

func TestTree_RemoveChildAt(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	node.AddChild(child)
	node.RemoveChildAt(0)

	assert.Equal(0, len(node.Children))
}

func TestTree_LenChildren(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	node.AddChild(child)

	assert.Equal(1, node.LenChildren())
}

func TestTree_SetParent(t *testing.T) {
	assert := assert.New(t)

	node := NewTreeNode("key", "value")
	child := NewTreeNode("child", "value")
	child.SetParent(node)

	assert.Equal(node, child.Parent)
}
