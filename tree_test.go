package gblink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Set(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	assert.Equal("one", tree.Root.Value)
}

func TestTree_Get(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	value, err := tree.Get(2)
	assert.Nil(err)
	assert.Equal("two", value)

	_, err = tree.Get(4)
	assert.NotNil(err)
}

func TestTree_Delete(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	tree.Delete(2)
	assert.Equal("one", tree.Root.Value)
	assert.Equal("three", tree.Root.Right.Value)
}

func TestTree_Len(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	assert.Equal(3, tree.Len())
}

func TestTree_DeleteMin(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	tree.DeleteMin(tree.Root)
	assert.Equal("one", tree.Root.Value)
	assert.Equal("two", tree.Root.Right.Value)
}

func TestTree_Min(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	assert.Equal("one", tree.Min(tree.Root).Value)
}

func TestTree_Keys(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	keys := tree.Keys()
	assert.Equal(3, len(keys))
	assert.Equal(1, keys[0])
	assert.Equal(2, keys[1])
	assert.Equal(3, keys[2])
}

func TestTree_Max(t *testing.T) {
	assert := assert.New(t)

	tree := NewTree[int, string]()
	tree.Set(1, "one")
	tree.Set(2, "two")
	tree.Set(3, "three")

	maxKey, err := tree.Max()
	assert.Nil(err)
	assert.Equal(3, maxKey)
}
